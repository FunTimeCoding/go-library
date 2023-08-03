package system

import (
	"bufio"
	"github.com/funtimecoding/go-library/errors"
)

func SaveFile(
	name string,
	text string,
) {
	f := Create(Join(WorkingDirectory(), name))

	w := bufio.NewWriter(f)
	_, e := w.WriteString(text)
	errors.PanicOnError(e)

	errors.PanicOnError(w.Flush())
	errors.PanicClose(f)
}
