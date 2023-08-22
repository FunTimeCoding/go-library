package errors

import "io"

func PanicClose(c io.Closer) {
	PanicOnError(c.Close())
}
