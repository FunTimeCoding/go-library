package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gorm.io/gorm"
)

type Store struct {
	database *gorm.DB
}

func (s *Store) Close() {
	d, e := s.database.DB()
	errors.PanicOnError(e)
	errors.PanicOnError(d.Close())
}
