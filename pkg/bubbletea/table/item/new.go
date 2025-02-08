package item

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/funtimecoding/go-library/pkg/strings"
)

func New() table.Model {
	columns := []table.Column{
		{Title: "Name", Width: 0},
	}
	rows := []table.Row{
		{strings.Alfa},
		{strings.Bravo},
		{strings.Charlie},
		{strings.Delta},
	}

	return table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)
}
