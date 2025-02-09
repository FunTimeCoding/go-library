package monitor

import "github.com/charmbracelet/bubbles/table"

type Model struct {
	table     *table.Model
	width     int
	height    int
	topBar    string
	bottomBar string
	second    int
}
