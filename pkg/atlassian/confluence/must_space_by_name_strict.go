package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/space"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustSpaceByNameStrict(name string) *space.Space {
	result, e := c.SpaceByNameStrict(name)
	errors.PanicOnError(e)

	return result
}
