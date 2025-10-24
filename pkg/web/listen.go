package web

import (
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func Listen(m *http.ServeMux) {
	ListenAddress(m, constant.Listen)
}
