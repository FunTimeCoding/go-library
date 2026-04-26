package store

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store/entry"
)

func (s *Store) Delete(id uint) error {
	if e := s.database.Delete(entry.New(), id).Error; e != nil {
		return fmt.Errorf("failed to delete entry: %w", e)
	}

	return nil
}
