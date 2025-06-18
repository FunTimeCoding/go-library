package bubbletea

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func Log(path string) *os.File {
	f, err := tea.LogToFile(path, "")
	errors.PanicOnError(err)

	return f
}
