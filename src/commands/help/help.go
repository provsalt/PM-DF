package help

import (
	"github.com/df-mc/dragonfly/server/cmd"
)

type Help struct{}

func (Help) Run(_ cmd.Source, o *cmd.Output) {
	for s, command := range cmd.Commands() {
		o.Printf("/%s: %s", s, command.Description())
	}
}
