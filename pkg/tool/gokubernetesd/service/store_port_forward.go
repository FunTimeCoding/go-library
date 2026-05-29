package service

func (s *Service) StorePortForward(
	identifier string,
	state interface{},
) {
	s.portForwards.Store(identifier, state)
}
