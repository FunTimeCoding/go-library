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
		s.database.Model(session.New()).
			Where("name = ?", name).
			Updates(
				map[string]any{
					"topic":            topic,
					"files":            files,
					"needs_reannounce": false,
				},
			).Error,
	)
	s.notify()
}
