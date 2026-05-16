package sublime

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/sublime/view"
)

func (c *Client) MustEditView(
	identifier int,
	old string,
	new string,
	all bool,
) *view.View {
	result, e := c.EditView(identifier, old, new, all)
	errors.PanicOnError(e)

	return result
}
