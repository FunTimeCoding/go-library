package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/space"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustSpaceByName(name string) *space.Space {
	result, e := c.SpaceByName(name)
	errors.PanicOnError(e)

	return result
}
