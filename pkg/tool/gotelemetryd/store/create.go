package store

import "github.com/funtimecoding/go-library/pkg/errors"

func (s *Store) Create(e *UsageEvent) {
	errors.PanicOnError(s.mapper.Create(e).Error)
}
