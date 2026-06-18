package server

import "net"

func (s *Server) WithListener(l net.Listener) *Server {
	s.listener = l

	return s
}
