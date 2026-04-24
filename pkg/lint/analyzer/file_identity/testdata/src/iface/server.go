package iface

import "net/http"

type Server struct{}

func (s *Server) handle(
	_ http.ResponseWriter,
	_ *http.Request,
) {
}
