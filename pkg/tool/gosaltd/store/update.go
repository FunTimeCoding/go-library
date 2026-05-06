package store

import "github.com/funtimecoding/go-library/pkg/errors"

func (s *Store) Update(r *HighstateRun) {
	errors.PanicOnError(s.mapper.Save(r).Error)
}
