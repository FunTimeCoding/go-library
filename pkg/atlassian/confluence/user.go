package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/user"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) User() *user.User {
	var result *response.User
	notation.DecodeStrict(
		c.basic.Get("/user/current"),
		&result,
		false,
	)

	return user.New(result)
}
