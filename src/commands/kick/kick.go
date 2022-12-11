package kick

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type Kick struct {
	Player []cmd.Target
	Reason cmd.Optional[string]
}

func (k Kick) Run(_ cmd.Source, o *cmd.Output) {
	for _, target := range k.Player {
		p, ok := target.(*player.Player)
		if !ok {
			o.Errorf("Usage: /kick <player> [reason ...]")
			return
		}

		if reason, ok2 := k.Reason.Load(); ok2 {
			p.Disconnect("Kicked by admin. Reason: " + reason)
			o.Printf("Kicked %s from the game: %s", p.Name(), reason)
			continue
		}
		p.Disconnect("Kicked by admin.")
		o.Printf("Kicked %s from the game", p.Name())
		continue
	}
}
