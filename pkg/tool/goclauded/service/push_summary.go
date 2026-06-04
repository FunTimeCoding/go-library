package service

func (s *Service) pushSummary(
	name string,
	body string,
	metadata map[string]string,
) error {
	return s.indexer.Push(name, body, metadata)
}
