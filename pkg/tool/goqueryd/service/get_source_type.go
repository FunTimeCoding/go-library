package service

func (s *Service) GetSourceType(
	collection string,
	pathPrefix string,
) string {
	return s.store.GetSourceType(collection, pathPrefix)
}
