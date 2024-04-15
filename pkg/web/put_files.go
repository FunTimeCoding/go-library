package web

import (
	"bytes"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/mime"
	"mime/multipart"
)

func PutFiles(
	locator string,
	headers map[string]string,
	files ...*mime.File,
) (int, string) {
	var b = new(bytes.Buffer)
	var w = multipart.NewWriter(b)

	for _, f := range files {
		mime.CreateAndWrite(w, f)
	}

	errors.PanicClose(w)
	request := NewPutBytes(locator, b)
	request.Header.Add(ContentTypeHeader, w.FormDataContentType())

	for k, v := range headers {
		request.Header.Add(k, v)
	}

	response := Send(Client(true), request)
	defer errors.PanicClose(response.Body)

	return response.StatusCode, ReadString(response)
}
