package service

func (s *Service) pushSummary(
	name string,
	body string,
	metadata map[string]string,
) error {
	return s.summaryIndexer.Push(name, body, metadata)
}
