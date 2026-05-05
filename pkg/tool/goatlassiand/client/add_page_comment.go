package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) AddPageComment(
	identifier string,
	body string,
) string {
	result, e := c.client.AddPageComment(
		c.context,
		identifier,
		client.CommentRequest{Body: body},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
