package system

import (
	"bufio"
	errors2 "github.com/funtimecoding/go-library/pkg/errors"
)

func SaveFile(
	path string,
	text string,
) {
	f := Create(path)

	w := bufio.NewWriter(f)
	_, e := w.WriteString(text)
	errors2.PanicOnError(e)

	errors2.PanicOnError(w.Flush())
	errors2.PanicClose(f)
}
