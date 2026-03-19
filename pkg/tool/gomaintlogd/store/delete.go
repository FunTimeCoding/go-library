package store

import "fmt"

func (s *Store) Delete(id uint) error {
	if e := s.database.Delete(&Entry{}, id).Error; e != nil {
		return fmt.Errorf("failed to delete entry: %w", e)
	}

	return nil
}
