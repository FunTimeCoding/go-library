package store

import "github.com/funtimecoding/go-library/pkg/errors"

func (s *Store) Create(r *TerraformRun) {
	errors.PanicOnError(s.mapper.Create(r).Error)
}
