package service

func (s *Service) SetPosition(
	identifier uint,
	target int,
) {
	s.store.SetPosition(identifier, target)
	s.notifier.Notify()
}
