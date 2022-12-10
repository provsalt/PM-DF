package config

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/df-mc/dragonfly/server"
	"github.com/pelletier/go-toml"
	"github.com/provsalt/PM-DF/src/wizard"
	"github.com/sirupsen/logrus"
	"os"
)

// Config is the server configuration file to be changed by the end user.
type Config struct {
	server.UserConfig

	//TODO Other configs
}

// ReadConfig reads the configuration from the config.toml file, or creates the file if it does not yet exist.
func ReadConfig(log *logrus.Logger) (server.Config, error) {
	c := server.DefaultConfig()
	var zero server.Config
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		p := tea.NewProgram(wizard.InitialModel())

		if err := p.Start(); err != nil {
			log.Fatal(err)
		}
		return zero, nil
	}
	data, err := os.ReadFile("config.toml")
	if err != nil {
		return zero, fmt.Errorf("error reading config: %v", err)
	}
	if err := toml.Unmarshal(data, &c); err != nil {
		return zero, fmt.Errorf("error decoding config: %v", err)
	}
	return c.Config(log)
}

// defaultConfig returns a configuration with the default values filled out.
func defaultConfig() Config {
	c := Config{}
	c.Network.Address = ":19132"
	c.Server.Name = "Dragonfly Server"
	c.Server.ShutdownMessage = "Server closed."
	c.Server.AuthEnabled = true
	c.Server.JoinMessage = "%v has joined the game"
	c.Server.QuitMessage = "%v has left the game"
	c.World.Folder = "world"
	c.Players.MaximumChunkRadius = 32
	c.Players.SaveData = true
	c.Players.Folder = "players"
	return c
}
