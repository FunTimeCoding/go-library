package scan

func (s *Service) checkOption(path string) {
	if !s.Option {
		s.addConcern(MissingOptionKey, MissingOptionText, path)
	}
}
