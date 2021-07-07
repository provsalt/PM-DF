package console

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Console struct{}

func (c Console) Name() string {
	return "CONSOLE"
}

func (c Console) Position() mgl64.Vec3 {
	return [3]float64{}
}

func (c Console) World() *world.World {
	return nil
}

func (Console) SendCommandOutput(output *cmd.Output) {
	for _, m := range output.Messages() {
		fmt.Println(m)
	}

	for _, e := range output.Errors() {
		fmt.Println(e.Error())
	}
}
