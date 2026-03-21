package web

import (
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (s *Server) isHTMX(r *http.Request) bool {
	return r.Header.Get(constant.ExtendedRequest) == "true"
}
