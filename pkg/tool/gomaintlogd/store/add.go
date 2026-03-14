package store

import "fmt"

func (s *Store) Add(v *Entry) error {
	if e := s.database.Create(v).Error; e != nil {
		return fmt.Errorf("failed to create entry: %w", e)
	}

	return nil
}
