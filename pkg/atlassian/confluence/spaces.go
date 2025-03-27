package confluence

import (
	kaos "github.com/essentialkaos/go-confluence/v6"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Spaces() []*kaos.Space {
	result, e := c.kaos.GetSpaces(kaos.SpaceParameters{})
	errors.PanicOnError(e)

	return result.Results
}
