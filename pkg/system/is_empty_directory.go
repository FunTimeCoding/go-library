package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"io"
)

func IsEmptyDirectory(path string) bool {
	f := Open(path)
	defer errors.LogClose(f)

	if _, e := f.Readdirnames(1); e == io.EOF {
		return true
	}

	return false
}
