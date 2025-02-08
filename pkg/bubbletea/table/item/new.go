package item

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/funtimecoding/go-library/pkg/bubbletea/style"
	"github.com/funtimecoding/go-library/pkg/strings"
)

func New() table.Model {
	columns := []table.Column{
		{Title: "Name", Width: 4},
	}
	rows := []table.Row{
		{strings.Alfa},
		{strings.Bravo},
		{strings.Charlie},
		{strings.Delta},
	}
	result := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)
	style.Table(result)

	return result
}
