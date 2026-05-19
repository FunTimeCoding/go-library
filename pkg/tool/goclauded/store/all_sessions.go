package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) AllSessions() []session.Session {
	var result []session.Session
	errors.PanicOnError(
		s.database.
			Order("last_seen DESC").
			Find(&result).Error,
	)

	return result
}
