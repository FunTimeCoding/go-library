package service

func (s *Service) DeleteDocument(
	collection string,
	path string,
) (bool, error) {
	return s.store.DeleteDocument(collection, path)
}
