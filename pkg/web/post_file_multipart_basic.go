package web

import (
	"bytes"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/mime"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"mime/multipart"
)

func PostFileMultipartBasic(
	locator string,
	user string,
	password string,
	fileType string,
	fileName string,
	b []byte,
	contentType string,
) (int, string) {
	var buffer = new(bytes.Buffer)
	var w = multipart.NewWriter(buffer)
	mime.CreateAndWrite(w, fileType, fileName, b)
	errors.PanicClose(w)
	request := NewPostBytes(locator, buffer)

	if contentType == "" {
		contentType = w.FormDataContentType()
	}

	request.Header.Add(constant.ContentTypeHeader, contentType)
	request.SetBasicAuth(user, password)

	for k, v := range request.Header {
		fmt.Printf("Request header: k=%s v=%s\n", k, join.Comma(v))
	}

	response := Send(Client(true), request)
	defer errors.PanicClose(response.Body)

	return response.StatusCode, ReadString(response)
}
