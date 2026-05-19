package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
	"time"
)

func (s *Store) sweepInactivity() {
	var sessions []session.Session
	cutoff := s.clock().Add(-1 * time.Hour)
	errors.PanicOnError(
		s.database.
			Where("last_seen < ? AND topic != '' AND timed_out = ''", cutoff).
			Find(&sessions).Error,
	)

	for _, e := range sessions {
		errors.PanicOnError(
			s.database.Model(session.New()).
				Where("name = ?", e.Name).
				Update("timed_out", constant.InactivityTimeout).Error,
		)
		s.LogEvent(
			e.Identifier,
			constant.InactivityTimeout,
			e.Name,
			e.Topic,
			"",
		)
	}

	if len(sessions) > 0 {
		s.notify()
	}
}
