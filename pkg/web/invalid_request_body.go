package web

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"net/http"
)

func InvalidRequestBody(w http.ResponseWriter) {
	http.Error(w, constant.InvalidRequestBody, http.StatusBadRequest)
}
