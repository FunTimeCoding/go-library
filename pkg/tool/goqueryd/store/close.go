package store

import "github.com/funtimecoding/go-library/pkg/errors"

func (s *Store) Close() {
	errors.PanicClose(s.database)
}
