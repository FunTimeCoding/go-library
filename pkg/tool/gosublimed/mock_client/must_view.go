package mock_client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/sublime/view"
)

func (c *Client) MustView(identifier int) *view.View {
	result, e := c.View(identifier)
	errors.PanicOnError(e)

	return result
}
