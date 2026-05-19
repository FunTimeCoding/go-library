package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/event"
)

func (s *Store) LogEvent(
	sessionIdentifier string,
	kind string,
	name string,
	target string,
	body string,
) {
	errors.PanicOnError(
		s.database.Create(
			event.New(
				sessionIdentifier,
				kind,
				name,
				target,
				body,
			),
		).Error,
	)
	s.notify()
}
