package service

func (s *Service) BindModelContextSession(
	name string,
	modelContextSessionIdentifier string,
) {
	s.store.BindModelContextSession(name, modelContextSessionIdentifier)
	s.notify()
}
