package confluence

import (
	kaos "github.com/essentialkaos/go-confluence/v6"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Groups() []*kaos.Group {
	result, e := c.kaos.GetGroups(kaos.CollectionParameters{})
	errors.PanicOnError(e)

	return result.Results
}
