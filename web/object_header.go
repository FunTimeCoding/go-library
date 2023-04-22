package web

import "net/http"

func ObjectHeader(w http.ResponseWriter) {
	w.Header().Set(ContentTypeHeader, ObjectContentType)
}
