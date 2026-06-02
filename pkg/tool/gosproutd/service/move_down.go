package service

func (s *Service) MoveDown(identifier uint) {
	s.store.MoveDown(identifier)
	s.notifier.Notify()
}
