package constant

import "github.com/charmbracelet/lipgloss"

var (
	Default = lipgloss.NewStyle()

	Table = lipgloss.NewStyle().BorderStyle(
		lipgloss.NormalBorder(),
	).BorderForeground(lipgloss.Color("240"))

	Modal = lipgloss.NewStyle().BorderStyle(
		lipgloss.NormalBorder(),
	).Width(50).Height(10).Align(lipgloss.Center)
)
