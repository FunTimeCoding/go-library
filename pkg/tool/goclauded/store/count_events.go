package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/event"
)

func (s *Store) CountEvents() int64 {
	var result int64
	errors.PanicOnError(
		s.database.Model(event.Stub()).Count(&result).Error,
	)

	return result
}
