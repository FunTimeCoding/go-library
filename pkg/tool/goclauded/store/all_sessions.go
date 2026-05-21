package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) AllSessions(
	limit int,
	offset int,
) []session.Session {
	var result []session.Session
	q := s.database.Order("last_seen DESC")

	if limit > 0 {
		q = q.Limit(limit)
	}

	if offset > 0 {
		q = q.Offset(offset)
	}

	errors.PanicOnError(q.Find(&result).Error)

	return result
}
