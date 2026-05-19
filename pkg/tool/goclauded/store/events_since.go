package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/event"
	"time"
)

func (s *Store) EventsSince(
	since time.Time,
	before time.Time,
	kind string,
	limit int,
	offset int,
) []event.Event {
	query := s.database.Model(event.Stub())

	if !since.IsZero() {
		query = query.Where("created_at >= ?", since)
	}

	if !before.IsZero() {
		query = query.Where("created_at <= ?", before)
	}

	if kind != "" {
		query = query.Where("kind = ?", kind)
	}

	var result []event.Event
	errors.PanicOnError(
		query.
			Order("created_at DESC").
			Limit(limit).
			Offset(offset).
			Find(&result).Error,
	)

	return result
}
