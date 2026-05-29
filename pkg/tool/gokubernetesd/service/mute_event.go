package service

func (s *Service) MuteEvent(
	reason string,
	message string,
	cluster string,
) error {
	return s.store.CreateMuteRule(reason, message, cluster)
}
