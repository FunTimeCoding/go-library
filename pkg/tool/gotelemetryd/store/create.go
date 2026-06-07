package store

func (s *Store) Create(e *UsageEvent) error {
	return s.mapper.Create(e).Error
}
