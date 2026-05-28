package service

func (s *Service) RemoveContext(
	collection string,
	pathPrefix string,
) bool {
	return s.store.RemoveContext(collection, pathPrefix)
}
