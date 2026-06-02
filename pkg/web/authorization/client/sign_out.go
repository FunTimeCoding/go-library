package client

import "net/http"

func (c *Client) SignOut(
	w http.ResponseWriter,
	r *http.Request,
) {
	http.SetCookie(
		w,
		&http.Cookie{
			Name:   subjectCookie,
			MaxAge: -1,
			Path:   "/",
		},
	)
	http.Redirect(w, r, "/", http.StatusFound)
}
