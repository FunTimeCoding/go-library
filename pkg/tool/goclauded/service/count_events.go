package service

func (s *Service) CountEvents() int64 {
	return s.store.CountEvents()
}
