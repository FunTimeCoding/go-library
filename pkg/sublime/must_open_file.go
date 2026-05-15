package sublime

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/sublime/view"
)

func (c *Client) MustOpenFile(path string) view.View {
	result, e := c.OpenFile(path)
	errors.PanicOnError(e)

	return result
}
