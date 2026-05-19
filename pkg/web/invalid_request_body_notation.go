package web

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"net/http"
)

func InvalidRequestBodyNotation(w http.ResponseWriter) {
	EncodeError(w, http.StatusBadRequest, constant.InvalidRequestBody)
}
