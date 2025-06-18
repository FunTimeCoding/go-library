package bubbletea

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func Run(
	m tea.Model,
	log bool,
) {
	if log {
		f := Log("tmp/tea.txt")
		defer errors.LogClose(f)
	}

	RunProgram(tea.NewProgram(m))
}
