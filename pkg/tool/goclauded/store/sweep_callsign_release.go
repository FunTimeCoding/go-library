package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
	"time"
)

func (s *Store) SweepCallsignRelease(cutoff time.Time) []session.Session {
	var sessions []session.Session
	errors.PanicOnError(
		s.database.
			Where(
				"callsign IS NOT NULL AND callsign != '' AND last_seen < ?",
				cutoff,
			).
			Find(&sessions).Error,
	)

	for _, e := range sessions {
		errors.PanicOnError(
			s.database.Model(session.Stub()).
				Where("callsign = ?", e.CallsignValue()).
				Update(constant.Callsign, nil).Error,
		)
	}

	return sessions
}
