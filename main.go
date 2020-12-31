package main

import (
	"fmt"
	"github.com/bradhe/stopwatch"
	"github.com/df-mc/dragonfly/dragonfly"
	"github.com/df-mc/dragonfly/dragonfly/player/chat"
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
	formatter "github.com/t-tomalak/logrus-easy-formatter"
	"io/ioutil"
	"os"
)

func main() {
	watch := stopwatch.Start()
	log := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &formatter.Formatter{
			TimestampFormat: "15:04:05",
			LogFormat:       "[%time%] [Server thread/%lvl%]: %msg% \n",
		},
	}
	log.Level = logrus.DebugLevel

	chat.Global.Subscribe(chat.StdoutSubscriber{})

	config, err := readConfig()
	if err != nil {
		log.Fatalln(err)
	}

	server := dragonfly.New(&config, log)
	server.CloseOnProgramEnd()
	if err := server.Start(); err != nil {
		log.Fatalln(err)
	}

	watch.Stop()
	log.Infof("Done (%v): For help, type \"help\" or \"?\" \n", watch.Milliseconds())

	for {
		if _, err := server.Accept(); err != nil {
			return
		}
	}
}

// readConfig reads the configuration from the config.toml file, or creates the file if it does not yet exist.
func readConfig() (dragonfly.Config, error) {
	c := dragonfly.DefaultConfig()
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return c, fmt.Errorf("failed encoding default config: %v", err)
		}
		if err := ioutil.WriteFile("config.toml", data, 0644); err != nil {
			return c, fmt.Errorf("failed creating config: %v", err)
		}
		return c, nil
	}
	data, err := ioutil.ReadFile("config.toml")
	if err != nil {
		return c, fmt.Errorf("error reading config: %v", err)
	}
	if err := toml.Unmarshal(data, &c); err != nil {
		return c, fmt.Errorf("error decoding config: %v", err)
	}
	return c, nil
}
