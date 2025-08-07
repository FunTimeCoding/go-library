package web

import (
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func TextHeader(w http.ResponseWriter) {
	w.Header().Set(constant.ContentTypeHeader, constant.TextContentType)
}
