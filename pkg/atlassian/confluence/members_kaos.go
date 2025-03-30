package confluence

import (
	kaos "github.com/essentialkaos/go-confluence/v6"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MembersKaos(group string) []*kaos.User {
	result, e := c.kaos.GetGroupMembers(group, kaos.CollectionParameters{})
	errors.PanicOnError(e)

	return result.Results
}
