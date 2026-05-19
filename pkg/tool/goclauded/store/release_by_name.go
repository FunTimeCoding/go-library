package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) ReleaseByName(name string) {
	errors.PanicOnError(
		s.database.Where("name = ?", name).Delete(session.New()).Error,
	)
	s.notify()
}
