package service

func (s *Service) AddContext(
	collection string,
	pathPrefix string,
	description string,
) {
	s.store.AddContext(collection, pathPrefix, description)
}
