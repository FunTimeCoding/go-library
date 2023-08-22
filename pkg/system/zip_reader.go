package system

import (
	"compress/gzip"
	"github.com/funtimecoding/go-library/pkg/errors"
	"io"
)

func ZipReader(r io.Reader) *gzip.Reader {
	result, e := gzip.NewReader(r)
	errors.PanicOnError(e)

	return result
}
