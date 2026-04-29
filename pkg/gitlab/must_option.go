package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MustOption() *gitlab.Settings {
	result, e := c.Option()
	errors.PanicOnError(e)

	return result
}
