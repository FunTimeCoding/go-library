package service

func (s *Service) notify() {
	s.notifier.Notify()
}
