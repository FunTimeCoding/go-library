package store

func (s *Store) Find(identifier uint) (*PlaybookRun, error) {
	var result PlaybookRun

	return &result, s.mapper.First(&result, identifier).Error
}
