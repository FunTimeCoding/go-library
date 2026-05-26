package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/pulse"
)

func (s *Store) SendPulse(
	sessionIdentifier string,
	fromName string,
	body string,
) {
	errors.PanicOnError(
		s.database.Create(
			pulse.New(sessionIdentifier, fromName, body),
		).Error,
	)
}
