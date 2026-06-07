package scan

func (s *Service) hasCapability() bool {
	return s.ModelContext ||
		s.Server ||
		s.Web ||
		s.Store ||
		s.Generated ||
		s.Convert ||
		s.Client ||
		s.Worker
}
