package client

import "net/http"

func (c *Client) Subject(r *http.Request) string {
	cookie, e := r.Cookie(subjectCookie)

	if e != nil {
		return ""
	}

	b, e := c.decrypt(cookie.Value)

	if e != nil {
		return ""
	}

	return string(b)
}
