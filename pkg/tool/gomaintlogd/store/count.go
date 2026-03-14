package store

func (s *Store) Count() int {
	var result int64
	s.database.Model(&Entry{}).Count(&result)

	return int(result)
}
