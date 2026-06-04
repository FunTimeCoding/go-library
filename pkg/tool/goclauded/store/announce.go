package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) Announce(
	name string,
	topic string,
	files string,
) {
	errors.PanicOnError(
		s.database.Model(session.Stub()).
			Where("callsign = ?", name).
			Updates(
				map[string]any{
					"topic":            topic,
					"files":            files,
					"needs_reannounce": false,
					"last_seen":        s.clock(),
				},
			).Error,
	)
}
