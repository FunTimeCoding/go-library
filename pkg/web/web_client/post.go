package web_client

import (
	"bytes"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
	"time"
)

func (c *Client) Post(
	locator string,
	body any,
) (*Response, error) {
	encoded := notation.Encode(body, false)
	start := c.clock.Now()
	response, e := http.Post(
		locator,
		web.ObjectContentType,
		bytes.NewBuffer([]byte(encoded)),
	)
	result := &Response{
		Response:     response,
		BodyBytes:    len(encoded),
		Milliseconds: time.Since(start).Milliseconds(),
	}

	return result, e
}
