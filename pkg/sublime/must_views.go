package sublime

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/sublime/view"
)

func (c *Client) MustViews() []*view.View {
	result, e := c.Views()
	errors.PanicOnError(e)

	return result
}
