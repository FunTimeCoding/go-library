package store

import "github.com/funtimecoding/go-library/pkg/errors"

func (s *Store) MustCount() int {
	result, e := s.Count()
	errors.PanicOnError(e)

	return result
}
