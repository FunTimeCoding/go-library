package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/alias"
)

func (s *Store) DeleteAlias(sessionIdentifier string) {
	errors.PanicOnError(
		s.database.Where(
			"session_identifier = ?",
			sessionIdentifier,
		).Delete(alias.Stub()).Error,
	)
	s.notify()
}
