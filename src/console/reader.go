package console

import (
	"bufio"
	"fmt"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"os"
	"strings"
	"time"
)

func StartConsole() {
	go func() {
		time.Sleep(time.Millisecond * 500)
		source := &Console{}
		// I don't use fmt.Scan() because the fmt package intentionally filters out whitespaces, this is how it is implemented.
		scanner := bufio.NewScanner(os.Stdin)
		for {
			if scanner.Scan() {
				commandString := scanner.Text()
				if commandString == "" {
					continue
				}
				commandName := strings.Split(commandString, " ")[0]
				command, ok := cmd.ByAlias(commandName)

				if !ok {
					output := &cmd.Output{}
					output.Errorf("Unknown command %v. Try /help for a list of commands", commandName)
					for _, e := range output.Errors() {
						fmt.Println(e)
					}
					continue
				}

				command.Execute(strings.TrimPrefix(strings.TrimPrefix(commandString, commandName), " "), source)
			}
		}
	}()
}
