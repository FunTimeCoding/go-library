package web

import (
	"bytes"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/mime"
	"mime/multipart"
)

func PostFiles(
	locator string,
	files ...*mime.File,
) []byte {
	var b = new(bytes.Buffer)
	var w = multipart.NewWriter(b)

	for _, f := range files {
		mime.CreateAndWrite(w, f)
	}

	errors.PanicClose(w)
	req := NewPostBytes(locator, b)
	req.Header.Add(ContentTypeHeader, w.FormDataContentType())
	c := Client(true)
	res := Send(c, req)
	defer errors.PanicClose(res.Body)

	return ReadBytes(res)
}
