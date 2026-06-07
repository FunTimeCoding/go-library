package service

func (s *Service) SendPulse(
	sessionIdentifier string,
	fromName string,
	body string,
) error {
	if e := s.store.SendPulse(sessionIdentifier, fromName, body); e != nil {
		return e
	}

	s.notify()

	return nil
}
