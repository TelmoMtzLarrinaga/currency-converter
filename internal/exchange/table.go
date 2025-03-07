package exchange

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("#c2f2d0")) // * #c2f2d0

var columns []table.Column = []table.Column{
	{Title: "COUNTRY", Width: 15},
	{Title: "CURRENCY", Width: 15},
	{Title: "UNIT", Width: 5},
}

var rows []table.Row = []table.Row{
	{"United States", "The Dollar", "USD"},
	{"Japan", "Japanese Yen", "JPY"},
	{"Spain", "Euro", "EUR"},
	{"England", "Pound Sterling", "GBP"},
}

// newTable returns a bubbletea table model with a custom format.
func newTable() table.Model {
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithHeight(5),
		table.WithFocused(true),
		table.WithKeyMap(table.DefaultKeyMap()),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#97ebdb")). // * #97ebdb
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#00c2c7")). // * #daf8e3
		Background(lipgloss.Color("#daf8e3")). // * #00c2c7
		Bold(false)
	t.SetStyles(s)

	return t
}
