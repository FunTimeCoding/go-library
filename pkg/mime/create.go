package mime

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"io"
	"mime/multipart"
)

func Create(
	w *multipart.Writer,
	fileType string,
	fileName string,
) io.Writer {
	result, e := w.CreateFormFile(fileType, fileName)
	errors.PanicOnError(e)

	return result
}
