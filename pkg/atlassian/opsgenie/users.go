package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/user"
	"github.com/funtimecoding/go-library/pkg/errors"
	raw "github.com/opsgenie/opsgenie-go-sdk-v2/user"
)

func (c *Client) Users() []*user.User {
	result, e := c.userClient.User.List(c.context, &raw.ListRequest{})
	errors.PanicOnError(e)

	return user.NewSlice(result.Users)
}
