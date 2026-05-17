package view

import (
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (v *View) IsExtendedRequest(r *http.Request) bool {
	return r.Header.Get(constant.ExtendedRequest) == "true"
}
