package service

func (s *Service) Send(
	name string,
	to string,
	body string,
) {
	s.store.SendMessage(name, to, body)
	s.notify()
}
