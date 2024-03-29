package system

import (
	"archive/zip"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func ZipReader(path string) *zip.ReadCloser {
	result, e := zip.OpenReader(path)
	errors.PanicOnError(e)

	return result
}
