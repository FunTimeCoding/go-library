package lifecycle

import (
	"github.com/funtimecoding/go-library/pkg/lifecycle/server"
	"net/http"
)

func WithServerMiddleware(
	address string,
	setup func(*http.ServeMux),
	middleware func(http.Handler) http.Handler,
) Option {
	return func(l *Lifecycle) {
		s := server.New(http.NewServeMux(), setup, address)
		s.Middleware = middleware
		l.component = append(l.component, s)
	}
}
