package model_context_client

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) Close() {
	errors.PanicClose(c.session)
}
