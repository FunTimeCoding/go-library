package basic

import "net/http"

func (c *Client) Post(
	path string,
	body any,
) ([]byte, error) {
	b, code, e := c.exchange(http.MethodPost, path, body)

	if e != nil {
		return nil, e
	}

	if code == http.StatusUnauthorized {
		c.login()
		b, code, e = c.exchange(http.MethodPost, path, body)

		if e != nil {
			return nil, e
		}
	}

	if code >= http.StatusBadRequest {
		return nil, parseDetail(b, code)
	}

	return b, nil
}
