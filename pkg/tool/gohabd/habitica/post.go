package habitica

import "net/http"

func (c *Client) post(
	path string,
	body any,
	result any,
) {
	r := c.do(http.MethodPost, path, body)
	c.decode(r, result)
}
