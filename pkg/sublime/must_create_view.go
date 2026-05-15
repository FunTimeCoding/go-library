package sublime

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/sublime/view"
)

func (c *Client) MustCreateView(
	title string,
	content string,
	syntax string,
) view.View {
	result, e := c.CreateView(title, content, syntax)
	errors.PanicOnError(e)

	return result
}
