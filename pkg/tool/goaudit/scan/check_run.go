package scan

func (s *Service) checkRun(path string) {
	if !s.Run {
		s.addConcern(MissingRunKey, MissingRunText, path)
	}
}
