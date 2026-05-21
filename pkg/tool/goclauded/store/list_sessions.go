package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
	"time"
)

func (s *Store) ListSessions() []session.Session {
	var result []session.Session
	cutoff := s.clock().Add(-1 * time.Hour)
	errors.PanicOnError(
		s.database.
			Where("last_seen > ? AND callsign IS NOT NULL", cutoff).
			Order(constant.Callsign).
			Find(&result).Error,
	)

	return result
}
