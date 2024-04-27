package web_client

import (
	"bytes"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/web_client/web_response"
	"net/http"
	"time"
)

func (c *Client) Post(
	locator string,
	body any,
) (*web_response.Response, error) {
	encoded := notation.Encode(body, false)
	start := c.clock.Now()
	response, e := http.Post(
		locator,
		web.ObjectContentType,
		bytes.NewBuffer([]byte(encoded)),
	)

	return &web_response.Response{
		Response:     response,
		BodyBytes:    len(encoded),
		Milliseconds: time.Since(start).Milliseconds(),
	}, e
}
