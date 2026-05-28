package service

func (s *Service) ReplaceTags(
	identifier int64,
	tags []string,
) error {
	return s.store.ReplaceTags(identifier, tags)
}
