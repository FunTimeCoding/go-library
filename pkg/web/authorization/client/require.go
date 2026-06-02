package client

import "net/http"

func (c *Client) Require(next http.HandlerFunc) http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		if c.Subject(r) == "" {
			http.Redirect(w, r, c.signInPath, http.StatusFound)

			return
		}

		next(w, r)
	}
}
