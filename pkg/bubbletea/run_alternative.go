package bubbletea

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func RunAlternative(m tea.Model) {
	p := tea.NewProgram(m, tea.WithAltScreen())
	_, e := p.Run()
	errors.PanicOnError(e)
}
