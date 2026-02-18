package lifecycle

import (
	"github.com/funtimecoding/go-library/pkg/lifecycle/server"
	"net/http"
)

func WithServer(
	address string,
	setup func(*http.ServeMux),
) Option {
	return func(l *Lifecycle) {
		l.component = append(
			l.component,
			&server.Server{
				Mux:     http.NewServeMux(),
				Setup:   setup,
				Address: address,
			},
		)
	}
}
