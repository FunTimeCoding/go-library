package service

func (s *Service) AddTags(
	identifier int64,
	tags []string,
) error {
	return s.store.AddTags(identifier, tags)
}
