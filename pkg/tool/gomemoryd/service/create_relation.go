package service

func (s *Service) CreateRelation(
	sourceIdentifier int64,
	targetIdentifier int64,
) error {
	return s.store.CreateRelation(sourceIdentifier, targetIdentifier)
}
