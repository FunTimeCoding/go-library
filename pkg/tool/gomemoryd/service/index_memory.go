package service

func (s *Service) IndexMemory(
	path string,
	content string,
	metadata map[string]string,
) error {
	return s.indexer.Push(path, content, metadata)
}
