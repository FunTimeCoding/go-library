package habitica

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/habitica/cron"
)

func (c *Client) MustCron() *cron.Cron {
	result, e := c.Cron()
	errors.PanicOnError(e)

	return result
}
