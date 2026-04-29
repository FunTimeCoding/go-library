package habitica

import "net/http"

func (c *Client) get(
	path string,
	result any,
) {
	r := c.do(http.MethodGet, path, nil)
	c.decode(r, result)
}
