package web

import "net/http"

func TextHeader(w http.ResponseWriter) {
	w.Header().Set(ContentTypeHeader, TextContentType)
}
