package store

import "github.com/funtimecoding/go-library/pkg/errors"

func (s *Store) Close() {
	errors.PanicOnError(s.database.Close())
}
