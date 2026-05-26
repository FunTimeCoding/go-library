package service

func (s *Service) SendPulse(
	sessionIdentifier string,
	fromName string,
	body string,
) {
	s.store.SendPulse(sessionIdentifier, fromName, body)
	s.notify()
}
