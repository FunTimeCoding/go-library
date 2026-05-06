package store

func (s *Store) Find(id uint) (*HighstateRun, error) {
	var result HighstateRun

	return &result, s.mapper.First(&result, id).Error
}
