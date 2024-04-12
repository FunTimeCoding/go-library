package mime

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"io"
	"mime/multipart"
	"path/filepath"
)

func Create(
	w *multipart.Writer,
	f *File,
) io.Writer {
	result, e := w.CreateFormFile(f.Type, filepath.Base(f.Name))
	errors.PanicOnError(e)

	return result
}
