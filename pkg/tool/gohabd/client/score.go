package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) Score(
	identifier string,
	direction string,
) string {
	result, e := c.client.ScoreTask(
		c.context,
		identifier,
		client.ScoreTaskParamsDirection(direction),
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
