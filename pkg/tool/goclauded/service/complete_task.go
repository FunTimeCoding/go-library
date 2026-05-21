package service

func (s *Service) CompleteTask(name string) string {
	topic := s.store.CompleteTask(name)

	if topic != "" {
		s.notify()
	}

	return topic
}
