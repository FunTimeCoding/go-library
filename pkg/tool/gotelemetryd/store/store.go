package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gorm.io/gorm"
)

type Store struct {
	mapper *gorm.DB
}

func (s *Store) Close() {
	d, e := s.mapper.DB()
	errors.PanicOnError(e)
	errors.PanicOnError(d.Close())
}
