package service

func (s *Service) Reorder(identifiers []uint) {
	s.store.Reorder(identifiers)
	s.notifier.Notify()
}
