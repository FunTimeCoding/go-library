package store

import "github.com/funtimecoding/go-library/pkg/errors"

func (s *Store) MustByName(name string) []Record {
	result, e := s.ByName(name)
	errors.PanicOnError(e)

	return result
}
