package store

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store/entry"
)

func (s *Store) Get(identifier uint) (*entry.Entry, error) {
	var result entry.Entry

	if e := s.database.First(&result, identifier).Error; e != nil {
		return nil, fmt.Errorf("failed to get entry: %w", e)
	}

	return &result, nil
}
