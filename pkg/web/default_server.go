package web

import (
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func DefaultServer(h http.Handler) *http.Server {
	return Server(h, constant.ListenAddress)
}
