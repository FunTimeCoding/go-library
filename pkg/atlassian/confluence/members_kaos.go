package confluence

import (
	kaos "github.com/essentialkaos/go-confluence/v6"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MembersKaos(group string) []*kaos.User {
	r, e := c.kaos.GetGroupMembers(group, kaos.CollectionParameters{})
	errors.PanicOnError(e)

	return r.Results
}
