package system

import (
	"compress/gzip"
	"github.com/funtimecoding/go-library/errors"
	"io"
)

func ZipReader(r io.Reader) *gzip.Reader {
	result, e := gzip.NewReader(r)
	errors.PanicOnError(e)

	return result
}
