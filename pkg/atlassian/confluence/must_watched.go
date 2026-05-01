package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustWatched() []*page.Page {
	result, e := c.Watched()
	errors.PanicOnError(e)

	return result
}
