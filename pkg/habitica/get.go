package habitica

import "net/http"

func (c *Client) get(
	path string,
	result any,
) error {
	r, e := c.do(http.MethodGet, path, nil)

	if e != nil {
		return e
	}

	return c.decode(r, result)
}
