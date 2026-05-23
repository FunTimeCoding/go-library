package service

func (s *Service) GetLabel(
	sessionIdentifier string,
	key string,
) string {
	return s.store.GetLabel(sessionIdentifier, key)
}
