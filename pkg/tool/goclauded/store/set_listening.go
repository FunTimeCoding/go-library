package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) SetListening(
	name string,
	listening bool,
) {
	errors.PanicOnError(
		s.database.Model(session.Stub()).
			Where("callsign = ?", name).
			Updates(
				map[string]any{
					"listening": listening,
					"last_seen": s.clock(),
				},
			).Error,
	)
}
