package lifecycle

import (
	"github.com/funtimecoding/go-library/pkg/lifecycle/server"
	"net/http"
	"time"
)

func WithServerTimeout(
	address string,
	setup func(*http.ServeMux),
	readTimeout time.Duration,
	writeTimeout time.Duration,
) Option {
	return func(l *Lifecycle) {
		s := server.New(http.NewServeMux(), setup, address)
		s.ReadTimeout = readTimeout
		s.WriteTimeout = writeTimeout
		l.component = append(l.component, s)
	}
}
