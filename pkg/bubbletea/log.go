package bubbletea

import (
	"charm.land/bubbletea/v2"
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func Log(path string) *os.File {
	result, e := tea.LogToFile(path, "")
	errors.PanicOnError(e)

	return result
}
