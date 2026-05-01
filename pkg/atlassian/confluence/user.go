package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/user"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) User() (*user.User, error) {
	r, e := c.basic.GetPath(constant.User)

	if e != nil {
		return nil, e
	}

	var result *response.User
	notation.DecodeStrict(r, &result, false)

	return user.New(result), nil
}
