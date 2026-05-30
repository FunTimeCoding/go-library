package service

func (s *Service) StopPortForward(identifier string) (any, bool) {
	return s.portForwards.LoadAndDelete(identifier)
}
