package exchange

import (
	"encoding/json"
	"net/http"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/zap"
)

// model will store our applications state.
type model struct {
	logger   *zap.Logger
	table    table.Model
	keys     keyMap
	help     help.Model
	quantity float64
	quote    table.Row
}

// Init is the first function that will be called. It can return an optional
// initial command or return nil if an initial command is not required.
func (m model) Init() tea.Cmd { return nil }

// Update is called when a message is received. The message will be inspected
// and, in response, the model will be updated and a command will be send.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can gracefully truncate
		// its view as needed.
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		// Quit currency-converter programn
		case "q":
			return m, tea.Quit
		// Help with key bindings
		case "?":
			m.help.ShowAll = !m.help.ShowAll
		// Show currency-converter programn
		case " ":
			// Parse obtained row for currency.
			m.quote = m.table.SelectedRow()
			quoteBytes := parse(m.logger, m.quote)

			// Make http request
			_, err := request(m.logger, m.quantity, quoteBytes) // TODO show value on table
			if err != http.StatusOK {
				return m, tea.Quit
			}
		}
	}
	m.table, cmd = m.table.Update(msg)
	// Return the updated model to the Bubble Tea runtime for processing.
	return m, cmd
}

// View renders the program's UI, which is just a string. The view is rendered
// after every Update.
func (m model) View() string {
	helpView := m.help.View(m.keys)
	return baseStyle.Render(m.table.View() + "\n\n" + helpView)
}

// We will define our applications initial state which is just a
func InitialModel(cfg *Config) model {
	return model{
		logger:   cfg.Logger,
		table:    newTable(),
		keys:     keys,
		help:     help.New(),
		quantity: cfg.Quantity,
	}
}

// parse parses the row into json encoding
func parse(logger *zap.Logger, row table.Row) []byte {
	if empty := len(row) > 0; !empty {
		logger.Error("exchange: row value is empty")
	}

	res, err := json.Marshal(row)
	if err != nil {
		logger.Error("exchange: couldn't marshal into JSON []byte: ", zap.Error(err))
	}

	return res
}
