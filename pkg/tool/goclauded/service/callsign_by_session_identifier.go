package service

func (s *Service) CallsignBySessionIdentifier(identifier string) string {
	return s.store.CallsignBySessionIdentifier(identifier)
}
