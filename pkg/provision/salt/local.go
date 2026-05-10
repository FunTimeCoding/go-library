package salt

import "github.com/funtimecoding/go-library/pkg/provision/salt/basic/response"

func (c *Client) Local(
	glob string,
	function string,
	a []string,
) (map[string]response.LocalReturn, error) {
	return c.basic.LocalClient(glob, function, a)
}
