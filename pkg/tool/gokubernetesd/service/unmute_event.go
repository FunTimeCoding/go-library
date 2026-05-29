package service

func (s *Service) UnmuteEvent(identifier uint) error {
	return s.store.DeleteMuteRule(identifier)
}
