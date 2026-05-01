package habitica

import "net/http"

func (c *Client) postDiscard(path string) error {
	r, e := c.do(http.MethodPost, path, nil)

	if e != nil {
		return e
	}

	return r.Body.Close()
}
