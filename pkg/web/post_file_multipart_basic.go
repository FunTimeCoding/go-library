package web

import (
	"bytes"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/mime"
	"mime/multipart"
)

func PostFileMultipartBasic(
	locator string,
	user string,
	password string,
	files ...*mime.File,
) (int, string) {
	var buffer = new(bytes.Buffer)
	var w = multipart.NewWriter(buffer)

	for _, f := range files {
		mime.CreateAndWrite(w, f)
	}

	errors.PanicClose(w)
	request := NewPostBytes(locator, buffer)
	request.Header.Add(ContentTypeHeader, w.FormDataContentType())
	request.SetBasicAuth(user, password)
	response := Send(Client(true), request)
	defer errors.PanicClose(response.Body)

	return response.StatusCode, ReadString(response)
}
