package server

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/location"
	"github.com/mark3labs/mcp-go/server"
	"log"
	"net/http"
)

func (s *Server) Serve() {
	m := http.NewServeMux()
	m.Handle(
		location.ModelContext,
		server.NewStreamableHTTPServer(
			s.server,
			server.WithLogger(s.Logger()),
		),
	)
	sse := server.NewSSEServer(s.server)
	m.Handle(location.Event, sse.SSEHandler())
	m.Handle(location.EventMessage, sse.MessageHandler())
	s.web = &http.Server{
		Addr:    Address,
		Handler: s.wrapAuthentication(m),
	}
	log.Printf("Listen %s\n", Address)
	web.ServeAsynchronous(s.web)
}
