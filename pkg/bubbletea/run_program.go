package bubbletea

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func RunProgram(p *tea.Program) {
	_, e := p.Run()
	errors.PanicOnError(e)
}
