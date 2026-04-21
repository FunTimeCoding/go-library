package web

import "net/http"

func EncodeNotation(
	w http.ResponseWriter,
	a any,
) {
	ObjectHeader(w)
	Encode(w, a)
}
