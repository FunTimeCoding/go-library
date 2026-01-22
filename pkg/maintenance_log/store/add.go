package store

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/maintenance_log/entry"
)

func (s *Store) Add(v *entry.Entry) error {
	if e := s.database.Create(v).Error; e != nil {
		return fmt.Errorf("failed to create entry: %w", e)
	}

	return nil
}
