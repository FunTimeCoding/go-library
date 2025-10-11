package confluence

import (
	kaos "github.com/essentialkaos/go-confluence/v6"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) SpacesKaos(keys []string) []*kaos.Space {
	r, e := c.kaos.GetSpaces(kaos.SpaceParameters{SpaceKey: keys})
	errors.PanicOnError(e)

	return r.Results
}
