package lifecycle

import (
	"github.com/funtimecoding/go-library/pkg/lifecycle/server"
	"net"
	"net/http"
)

func WithListener(
	l net.Listener,
	setup func(*http.ServeMux),
) Option {
	return func(c *Lifecycle) {
		c.component = append(
			c.component,
			server.NewWithListener(http.NewServeMux(), setup, l),
		)
	}
}
