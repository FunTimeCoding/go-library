package habitica

import (
	"github.com/funtimecoding/go-library/pkg/habitica/response"
	"net/http"
)

func (c *Client) user() (response.User, error) {
	var result response.User
	r, e := c.do(http.MethodGet, "/user", nil)

	if e != nil {
		return result, e
	}

	return result, c.decode(r, &result)
}
