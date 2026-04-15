package atl

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) TransitionIssue(key string, transitionIdentifier string) string {
	result, e := c.client.TransitionIssue(
		c.context,
		key,
		client.TransitionRequest{TransitionIdentifier: transitionIdentifier},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
