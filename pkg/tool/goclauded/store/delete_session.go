package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) DeleteSession(identifier string) {
	errors.PanicOnError(
		s.database.Where(
			"identifier = ?",
			identifier,
		).Delete(session.Stub()).Error,
	)
}
