package lifecycle

import (
	"github.com/funtimecoding/go-library/pkg/lifecycle/server"
	"net/http"
)

func WithProtectedServer(
	address string,
	setup func(*http.ServeMux),
) Option {
	return func(l *Lifecycle) {
		l.component = append(
			l.component,
			server.NewProtected(http.NewServeMux(), setup, address),
		)
	}
}
