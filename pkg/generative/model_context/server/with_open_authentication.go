package server

func WithOpenAuthentication(
	serverLocator string,
	authorizationLocator string,
	clientIdentifier string,
) Option {
	return func(s *Server) {
		s.openAuthentication = true
		s.serverLocator = serverLocator
		s.authorizationLocator = authorizationLocator
		s.clientIdentifier = clientIdentifier
	}
}
