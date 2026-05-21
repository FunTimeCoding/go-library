package service

func (s *Service) SetListening(
	name string,
	listening bool,
) {
	s.store.SetListening(name, listening)
	s.notify()
}
