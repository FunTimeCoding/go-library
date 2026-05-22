package basic

import "net/http"

func (c *Client) Get(path string) ([]byte, error) {
	b, code, e := c.exchange(http.MethodGet, path, nil)

	if e != nil {
		return nil, e
	}

	if code == http.StatusUnauthorized {
		c.login()
		b, code, e = c.exchange(http.MethodGet, path, nil)

		if e != nil {
			return nil, e
		}
	}

	if code >= http.StatusBadRequest {
		return nil, parseDetail(b, code)
	}

	return b, nil
}
