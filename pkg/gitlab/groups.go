package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

func (c *Client) Groups() []*gitlab.Group {
	result, _, e := c.client.Groups.ListGroups(nil, nil)
	errors.PanicOnError(e)

	return result
}
