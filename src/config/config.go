package config

import (
	"fmt"
	"github.com/df-mc/dragonfly/server"
	"github.com/pelletier/go-toml"
	"os"
)

// Config is the server configuration file to be changed by the end user.
type Config struct {
	server.Config

	//TODO Other configs
}

// ReadConfig reads the configuration from the config.toml file, or creates the file if it does not yet exist.
func ReadConfig() (server.Config, error) {
	c := server.DefaultConfig()
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return c, fmt.Errorf("failed encoding default config: %v", err)
		}
		if err := os.WriteFile("config.toml", data, 0644); err != nil {
			return c, fmt.Errorf("failed creating config: %v", err)
		}
		return c, nil
	}
	data, err := os.ReadFile("config.toml")
	if err != nil {
		return c, fmt.Errorf("error reading config: %v", err)
	}
	if err := toml.Unmarshal(data, &c); err != nil {
		return c, fmt.Errorf("error decoding config: %v", err)
	}
	return c, nil
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
	c.World.Name = "World"
	c.World.Folder = "world"
	c.World.SimulationDistance = 8
	c.Players.MaximumChunkRadius = 32
	c.Players.SaveData = true
	c.Players.Folder = "players"
	return c
}
