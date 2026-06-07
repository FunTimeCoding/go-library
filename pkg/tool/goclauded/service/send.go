package service

func (s *Service) Send(
	name string,
	to string,
	body string,
) error {
	if e := s.store.SendMessage(name, to, body); e != nil {
		return e
	}

	s.notify()

	return nil
}
