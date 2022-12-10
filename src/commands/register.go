package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/provsalt/PM-DF/src/commands/version"
)

// Register registers all commands in the commands directory
func Register() {
	cmd.Register(cmd.New("version", "Check the server version", []string{"ver"}, version.Version{}))
}
