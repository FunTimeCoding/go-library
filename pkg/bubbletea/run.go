package bubbletea

import tea "github.com/charmbracelet/bubbletea"

func Run(m tea.Model) {
	RunProgram(tea.NewProgram(m))
}
