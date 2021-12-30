package wizard

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/df-mc/dragonfly/server"
	"github.com/pelletier/go-toml"
	"os"
)

type model struct {
	textInput textinput.Model
	err       error
}
type tickMsg struct{}
type errMsg error

func InitialModel() model {
	ti := textinput.NewModel()
	ti.Placeholder = "y/N"
	ti.Focus()
	ti.CharLimit = 1
	ti.Width = 20

	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter:
			if m.textInput.Value() == "y" {
				data, err := toml.Marshal(server.DefaultConfig())
				if err != nil {
					panic("failed encoding default config")
				}
				if err := os.WriteFile("config.toml", data, 0644); err != nil {
					panic("failed creating config")
				}
				return m, tea.Quit
			}else {
				return m, tea.Quit
			}
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"Do you wish to skip the setup wizard? Setup wizard is unfinished. %s",
		m.textInput.View(),
	) + "\n"
}
