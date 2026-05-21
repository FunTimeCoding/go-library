package service

func (s *Service) pushSummary(
	name string,
	body string,
) error {
	return s.indexer.Push(name, body)
}
