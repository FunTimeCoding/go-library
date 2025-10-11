package confluence

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/virtomize/confluence-go-api"
)

func (c *Client) SpacesVirtomize() []goconfluence.Space {
	r, e := c.virtomize.GetAllSpaces(goconfluence.AllSpacesQuery{})
	errors.PanicOnError(e)

	return r.Results
}
