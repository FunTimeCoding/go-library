package service

func (s *Service) SetParent(
	identifier int64,
	parentIdentifier *int64,
) error {
	return s.store.SetParent(identifier, parentIdentifier)
}
