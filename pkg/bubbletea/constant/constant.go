package constant

import "github.com/charmbracelet/lipgloss"

var (
	Default = lipgloss.NewStyle()

	Table = lipgloss.NewStyle().BorderStyle(
		lipgloss.NormalBorder(),
	).BorderForeground(lipgloss.Color("240"))
)
