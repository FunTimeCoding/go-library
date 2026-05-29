package service

func (s *Service) StopPortForward(identifier string) (interface{}, bool) {
	return s.portForwards.LoadAndDelete(identifier)
}
