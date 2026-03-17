package store

import "fmt"

func (s *Store) Get(id uint) (*Entry, error) {
	var result Entry

	if e := s.database.First(&result, id).Error; e != nil {
		return nil, fmt.Errorf("failed to get entry: %w", e)
	}

	return &result, nil
}
