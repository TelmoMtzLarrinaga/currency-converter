package exchange

import (
	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/zap"
)

// model will store our applications state.
type model struct {
	logger *zap.Logger
}

// Init is the first function that will be called. It can return an optional
// initial command or return nil if an initial command is not required.
func (m model) Init() tea.Cmd {
	// No I/O as of now
	return nil
}

// Update is called when a message is received. The message will be inspected
// and, in response, the model will be updated and a command will be send.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Quit currency-converter programn
		case "q":
			return m, tea.Quit
		// Show currency-converter programn
		case "enter", " ":
		}
	}
	// Return the updated model to the Bubble Tea runtime for processing.
	return m, nil
}

// View renders the program's UI, which is just a string. The view is rendered
// after every Update.
func (m model) View() string {
	return ""
}

// We will define our applications initial state which is just a
func InitialModel(cfg *Config) model {
	return model{
		logger: cfg.Logger,
	}
}
