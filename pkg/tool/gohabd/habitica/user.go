package habitica

import (
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/habitica/response"
	"net/http"
)

func (c *Client) user() response.User {
	var result response.User
	r := c.do(http.MethodGet, "/user", nil)
	c.decode(r, &result)

	return result
}
