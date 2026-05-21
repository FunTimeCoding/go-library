package service

func (s *Service) ResolveModelContextSession(
	modelContextSessionIdentifier string,
) (string, string) {
	return s.store.ResolveModelContextSession(modelContextSessionIdentifier)
}
