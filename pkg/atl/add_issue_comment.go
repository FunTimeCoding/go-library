package atl

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) AddIssueComment(
	key string,
	body string,
) string {
	result, e := c.client.AddIssueComment(
		c.context,
		key,
		client.CommentRequest{Body: body},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
