package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/pulse"
)

func (s *Store) PulsesBySession(sessionIdentifier string) []pulse.Pulse {
	var result []pulse.Pulse
	errors.PanicOnError(
		s.database.
			Where("session_identifier = ?", sessionIdentifier).
			Order("created_at").
			Find(&result).Error,
	)

	return result
}
