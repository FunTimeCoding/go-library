package store

import "github.com/funtimecoding/go-library/pkg/errors"

func (s *Store) MustSave(r Record) string {
	result, e := s.Save(r)
	errors.PanicOnError(e)

	return result
}
