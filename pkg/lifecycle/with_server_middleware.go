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
		l.component = append(
			l.component,
			&server.Server{
				Mux:        http.NewServeMux(),
				Setup:      setup,
				Middleware: middleware,
				Address:    address,
			},
		)
	}
}
