package service

func (s *Service) CreateImpression(
	content string,
	source string,
) (int64, error) {
	return s.store.CreateImpression(content, source)
}
