package service

func (s *Service) StorePortForward(
	identifier string,
	state any,
) {
	s.portForwards.Store(identifier, state)
}
