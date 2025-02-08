package bubbletea

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func Run(m tea.Model) {
	p := tea.NewProgram(m)
	_, e := p.Run()
	errors.PanicOnError(e)
}
