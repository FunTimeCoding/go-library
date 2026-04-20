package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net"
	"net/http"
)

func ServeListenerAsynchronous(
	s *http.Server,
	l net.Listener,
) {
	go func() {
		errors.PanicOnUnclean(s.Serve(l))
	}()
}
