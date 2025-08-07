package web

import (
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func ListenAsynchronous(m *http.ServeMux) *http.Server {
	s := &http.Server{Addr: constant.ListenAddress, Handler: m}
	ServeAsynchronous(s)

	return s
}
