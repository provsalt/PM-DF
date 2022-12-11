package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/provsalt/PM-DF/src/commands/kick"
	"github.com/provsalt/PM-DF/src/commands/version"
)

// Register registers all commands in the commands directory
func Register() {
	cmd.Register(cmd.New("version", "Gets the version of this server", []string{"ver"}, version.Version{}))
	cmd.Register(cmd.New("kick", "Removes the specified player from the server", []string{"kick"}, kick.Kick{}))
}
