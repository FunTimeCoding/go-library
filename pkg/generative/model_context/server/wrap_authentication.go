package server

import (
	authorization "github.com/funtimecoding/go-library/pkg/web/authorization/constant"
	"github.com/funtimecoding/go-library/pkg/web/location"
	"net/http"
)

func (s *Server) wrapAuthentication(h http.Handler) http.Handler {
	m := http.NewServeMux()
	var middle func(http.Handler) http.Handler

	if s.tokenAuthentication || s.openAuthentication {
		middle = func(next http.Handler) http.Handler {
			return crossOriginMiddleware(s.authenticate(next))
		}
	} else {
		middle = crossOriginMiddleware
	}

	m.Handle(location.ModelContext, middle(h))
	m.Handle(location.Event, middle(h))
	m.Handle(location.EventMessage, middle(h))

	if s.openAuthentication {
		m.HandleFunc(authorization.ProtectedResource, s.protectedResource)
	}

	return m
}
