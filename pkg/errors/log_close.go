package errors

import "io"

func LogClose(c io.Closer) {
	LogOnError(c.Close())
}
