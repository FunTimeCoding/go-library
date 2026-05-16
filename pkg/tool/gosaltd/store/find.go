package store

func (s *Store) Find(identifier uint) (*HighstateRun, error) {
	var result HighstateRun

	return &result, s.mapper.First(&result, identifier).Error
}
