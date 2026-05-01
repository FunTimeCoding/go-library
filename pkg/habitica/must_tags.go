package habitica

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/habitica/response"
)

func (c *Client) MustTags() []response.Tag {
	result, e := c.Tags()
	errors.PanicOnError(e)

	return result
}
