package service

func (s *Service) FindSimilarFiles(
	path string,
	limit int,
) ([]string, error) {
	return s.store.FindSimilarFiles(path, limit)
}
