package monitor

import "github.com/charmbracelet/bubbles/table"

func New(m table.Model) *Model {
	return &Model{
		topBar:    "top bar",
		table:     m,
		bottomBar: "bottom bar",
	}
}
