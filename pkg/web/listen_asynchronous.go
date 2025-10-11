package web

import (
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func ListenAsynchronous(m *http.ServeMux) *http.Server {
	s := Server(m, constant.Listen)
	ServeAsynchronous(s)

	return s
}
