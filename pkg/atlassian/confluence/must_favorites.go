package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustFavorites() []*page.Page {
	result, e := c.Favorites()
	errors.PanicOnError(e)

	return result
}
