package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) UpdateTopic(
	name string,
	topic string,
	files string,
) {
	errors.PanicOnError(
		s.database.Model(session.Stub()).
			Where("callsign = ?", name).
			Updates(
				map[string]any{
					"topic":     topic,
					"files":     files,
					"last_seen": s.clock(),
				},
			).Error,
	)
}
