package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/pulse"
)

func (s *Store) PendingPulses(sessionIdentifier string) []pulse.Pulse {
	var result []pulse.Pulse
	errors.PanicOnError(
		s.database.
			Where(
				"session_identifier = ? AND from_name = '' AND consumed = ?",
				sessionIdentifier,
				false,
			).
			Order("created_at").
			Find(&result).Error,
	)

	if len(result) > 0 {
		errors.PanicOnError(
			s.database.Model(pulse.Stub()).
				Where(
					"session_identifier = ? AND from_name = '' AND consumed = ?",
					sessionIdentifier,
					false,
				).
				Update("consumed", true).Error,
		)
	}

	return result
}
