package service

func (s *Service) IsMuted(
	reason string,
	message string,
	cluster string,
) (bool, error) {
	return s.store.IsMuted(reason, message, cluster)
}
