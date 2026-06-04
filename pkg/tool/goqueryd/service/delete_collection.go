package service

func (s *Service) DeleteCollection(name string) bool {
	return s.store.DeleteCollection(name)
}
