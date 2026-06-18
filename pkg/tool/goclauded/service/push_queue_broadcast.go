package service

func (s *Service) PushQueueBroadcast(
	kind string,
	body string,
) error {
	sessions, e := s.store.ListSessions()

	if e != nil {
		return e
	}

	var callsigns []string

	for _, session := range sessions {
		callsigns = append(callsigns, session.CallsignValue())
	}

	if e := s.store.PushQueueBroadcast(callsigns, kind, body); e != nil {
		return e
	}

	s.notify()

	return nil
}
