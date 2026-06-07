package service

func (s *Service) SetListening(
	name string,
	listening bool,
) error {
	if e := s.store.SetListening(name, listening); e != nil {
		return e
	}

	s.notify()

	return nil
}
