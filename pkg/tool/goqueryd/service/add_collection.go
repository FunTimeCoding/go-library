package service

func (s *Service) AddCollection(
	name string,
	path string,
	pattern string,
) {
	s.store.AddCollection(name, path, pattern)
}
