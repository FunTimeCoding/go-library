package server

func (s *Server) WithCertificate(
	certificate string,
	key string,
) *Server {
	s.certificate = certificate
	s.key = key

	return s
}
