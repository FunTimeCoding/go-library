package service

func (s *Service) RemoveTags(
	identifier int64,
	tags []string,
) error {
	return s.store.RemoveTags(identifier, tags)
}
