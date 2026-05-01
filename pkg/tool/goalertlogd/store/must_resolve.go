package store

import "github.com/funtimecoding/go-library/pkg/errors"

func (s *Store) MustResolve(key string) {
	errors.PanicOnError(s.Resolve(key))
}
