package main

import (
	"github.com/bradhe/stopwatch"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/provsalt/PM-DF/src/commands"
	config2 "github.com/provsalt/PM-DF/src/config"
	"github.com/provsalt/PM-DF/src/console"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/sirupsen/logrus"
	formatter "github.com/t-tomalak/logrus-easy-formatter"
	"os"
)

func main() {
	watch := stopwatch.Start()
	log := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &formatter.Formatter{
			TimestampFormat: "15:04:05",
			LogFormat:       "[%time%] [Server/%lvl%]: %msg% \n",
		},
	}
	log.Level = logrus.DebugLevel

	chat.Global.Subscribe(chat.StdoutSubscriber{})

	config, err := config2.ReadConfig(log)
	if err != nil {
		log.Fatalln(err)
	}

	srv := config.New()
	srv.CloseOnProgramEnd()

	srv.Listen()

	if !config.AuthDisabled {
		log.Info("Online mode is enabled. The server will verify that players are authenticated to Xbox Live.")
	} else {
		log.Warn(text.ANSI(text.Colourf("<yellow>Online mode is disabled. The server will not verify that players are authenticated to Xbox Live.</yellow>")))
	}

	watch.Stop()
	log.Infof("Done (%v): For help, type \"help\" or \"?\" \n", watch.Milliseconds())
	console.StartConsole()
	commands.Register()

	for srv.Accept(func(p *player.Player) {
		log.Infof(text.ANSI(text.Colourf("<aqua>%s</aqua> [/%s] logged in with entity %d at (%s, %d, %d, %d)", p.Name(), p.Addr().String(), len(srv.World().Entities()), p.World().Name(), int(p.Position().X()), int(p.Position().Y()), int(p.Position().Z()))))
	}) {
	}
}
