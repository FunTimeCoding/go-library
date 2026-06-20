package store

import "github.com/funtimecoding/go-library/pkg/errors"

func (s *Store) Update(r *Run) {
	errors.PanicOnError(s.mapper.Table(s.tableName).Save(r).Error)
}
