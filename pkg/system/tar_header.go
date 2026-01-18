package system

import (
	"archive/tar"
	"github.com/funtimecoding/go-library/pkg/errors"
	"io/fs"
)

func TarHeader(
	i fs.FileInfo,
	link string,
) *tar.Header {
	result, e := tar.FileInfoHeader(i, link)
	errors.PanicOnError(e)

	return result
}
