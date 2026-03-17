package store

import "fmt"

func (s *Store) Update(v *Entry) error {
	if e := s.database.Save(v).Error; e != nil {
		return fmt.Errorf("failed to update entry: %w", e)
	}

	return nil
}
