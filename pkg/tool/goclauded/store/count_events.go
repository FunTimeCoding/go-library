package store

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/event"

func (s *Store) CountEvents() (int64, error) {
	var result int64

	return result, s.database.Model(event.Stub()).Count(&result).Error
}
