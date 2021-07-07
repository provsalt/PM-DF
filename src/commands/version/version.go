package version

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"runtime/debug"
)

// Version is the command that tells the player the version of the running server
type Version struct {
}

// Run ...
func (Version) Run(_ cmd.Source, output *cmd.Output) {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		output.Errorf("Failed to read build info")
		return
	}
	var dfVer string
	var gopherVer string
	for _, dep := range bi.Deps {
		if dep.Path == "github.com/df-mc/dragonfly" {
			dfVer = dep.Version
		}
		if dep.Path == "github.com/sandertv/gophertunnel" {
			gopherVer = dep.Version
		}
	}

	output.Printf("This server is running Dragonfly for Minecraft: Bedrock Edition v%s. Implementing: \nDragonfly: %s \nGophertunnel: %s", protocol.CurrentVersion, dfVer, gopherVer)
}
