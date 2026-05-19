package service

func (s *Service) ForwardImpression(
	content string,
	source string,
) {
	s.memory.SaveImpression(content, source)
}
