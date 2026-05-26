package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/pulse"
)

func (s *Store) LatestPulse(sessionIdentifier string) *pulse.Pulse {
	var result pulse.Pulse
	r := s.database.
		Where("session_identifier = ?", sessionIdentifier).
		Order("created_at DESC").
		Limit(1).
		Find(&result)
	errors.PanicOnError(r.Error)

	if r.RowsAffected == 0 {
		return nil
	}

	return &result
}
