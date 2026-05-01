package store

import "github.com/funtimecoding/go-library/pkg/errors"

func (s *Store) MustUnresolved() []UnresolvedRecord {
	result, e := s.Unresolved()
	errors.PanicOnError(e)

	return result
}
