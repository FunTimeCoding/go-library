package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) CreateDiscovered(identifier string) {
	errors.PanicOnError(
		s.database.Create(session.NewDiscovered(identifier, s.clock())).Error,
	)
}
