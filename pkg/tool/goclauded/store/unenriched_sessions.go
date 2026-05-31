package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) UnenrichedSessions() []session.Session {
	var result []session.Session
	errors.PanicOnError(
		s.database.
			Where("slug = '' OR slug IS NULL OR callsign IS NOT NULL").
			Find(&result).Error,
	)

	return result
}
