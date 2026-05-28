package service

func (s *Service) SetSourceType(
	collection string,
	pathPrefix string,
	sourceType string,
) {
	s.store.SetSourceType(collection, pathPrefix, sourceType)
}
