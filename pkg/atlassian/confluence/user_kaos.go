package confluence

import (
	kaos "github.com/essentialkaos/go-confluence/v6"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) UserKaos() *kaos.User {
	result, e := c.kaos.GetCurrentUser(kaos.ExpandParameters{})
	errors.PanicOnError(e)

	return result
}
