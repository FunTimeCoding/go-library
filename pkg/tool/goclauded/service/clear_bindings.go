package service

func (s *Service) ClearBindings() {
	s.store.ClearBindings()
}
