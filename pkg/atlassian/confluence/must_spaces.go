package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/space"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustSpaces() []*space.Space {
	result, e := c.Spaces()
	errors.PanicOnError(e)

	return result
}
