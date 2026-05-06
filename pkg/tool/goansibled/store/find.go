package store

func (s *Store) Find(id uint) (*PlaybookRun, error) {
	var result PlaybookRun

	return &result, s.mapper.First(&result, id).Error
}
