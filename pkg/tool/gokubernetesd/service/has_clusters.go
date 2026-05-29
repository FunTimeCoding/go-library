package service

func (s *Service) HasClusters() bool {
	return len(s.clusters) > 0
}
