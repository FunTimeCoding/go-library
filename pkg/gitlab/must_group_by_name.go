package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MustGroupByName(s string) *gitlab.Group {
	result, e := c.GroupByName(s)
	errors.PanicOnError(e)

	return result
}
