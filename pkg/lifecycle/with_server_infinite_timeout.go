package lifecycle

import (
	"github.com/funtimecoding/go-library/pkg/lifecycle/server"
	"net/http"
)

func WithServerInfiniteTimeout(
	address string,
	setup func(*http.ServeMux),
) Option {
	return func(l *Lifecycle) {
		s := server.New(http.NewServeMux(), setup, address)
		s.ReadTimeout = -1
		s.WriteTimeout = -1
		l.component = append(l.component, s)
	}
}
