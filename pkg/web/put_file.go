package web

import (
	"bytes"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func PutFile(
	locator string,
	headers map[string]string,
	b []byte,
) (int, string) {
	var buffer = new(bytes.Buffer)
	buffer.Write(b)
	request := NewPutBytes(locator, buffer)

	for k, v := range headers {
		request.Header.Add(k, v)
	}

	response := Send(Client(true), request)
	defer errors.PanicClose(response.Body)

	return response.StatusCode, ReadString(response)
}
