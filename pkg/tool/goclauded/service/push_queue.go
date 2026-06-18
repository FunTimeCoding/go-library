package service

func (s *Service) PushQueue(
	callsign string,
	kind string,
	body string,
) error {
	if e := s.store.PushQueue(callsign, kind, body); e != nil {
		return e
	}

	s.notify()

	return nil
}
