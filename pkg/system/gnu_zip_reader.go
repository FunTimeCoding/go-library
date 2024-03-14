package system

import (
	"compress/gzip"
	"github.com/funtimecoding/go-library/pkg/errors"
	"io"
)

func GnuZipReader(r io.Reader) *gzip.Reader {
	result, e := gzip.NewReader(r)
	errors.PanicOnError(e)

	return result
}
