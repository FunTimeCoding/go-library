package web

import (
	"net/http"
	"time"
)

func UnsetCookie(
	w http.ResponseWriter,
	k string,
) {
	http.SetCookie(
		w,
		&http.Cookie{
			Name:     k,
			Value:    "",
			HttpOnly: true,
			Secure:   true,
			Expires:  time.Unix(0, 0),
		},
	)
}
