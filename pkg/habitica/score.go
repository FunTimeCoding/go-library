package habitica

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/client"
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
