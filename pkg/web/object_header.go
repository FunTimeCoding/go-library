package web

import (
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func ObjectHeader(w http.ResponseWriter) {
	w.Header().Set(constant.ContentType, constant.Object)
}
