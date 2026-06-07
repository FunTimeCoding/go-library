package store

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/event"
	"time"
)

func (s *Store) EventsSince(
	since time.Time,
	before time.Time,
	kind string,
	limit int,
	offset int,
) ([]event.Event, error) {
	q := s.database.Model(event.Stub())

	if !since.IsZero() {
		q = q.Where("created_at >= ?", since)
	}

	if !before.IsZero() {
		q = q.Where("created_at <= ?", before)
	}

	if kind != "" {
		q = q.Where("kind = ?", kind)
	}

	var result []event.Event

	return result, q.Order(
		"created_at DESC",
	).Limit(limit).Offset(offset).Find(&result).Error
}
