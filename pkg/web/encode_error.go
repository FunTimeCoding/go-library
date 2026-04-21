package web

import "net/http"

func EncodeError(
	w http.ResponseWriter,
	code int,
	message string,
) {
	ObjectHeader(w)
	w.WriteHeader(code)
	Encode(w, map[string]string{"error": message})
}
