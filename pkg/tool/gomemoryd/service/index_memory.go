package service

func (s *Service) IndexMemory(
	path string,
	content string,
) error {
	return s.indexer.Push(path, content)
}
