package habitica

import "net/http"

func (c *Client) post(
	path string,
	body any,
	result any,
) error {
	r, e := c.do(http.MethodPost, path, body)

	if e != nil {
		return e
	}

	return c.decode(r, result)
}
