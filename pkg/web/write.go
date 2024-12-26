package web

import "net/http"

func Write(
	w http.ResponseWriter,
	code int,
	s string,
) {
	WriteBytes(w, code, []byte(s))
}
