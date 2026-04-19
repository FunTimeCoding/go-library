package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/raid"
)

func (s *Store) Fights() []raid.Fight {
	var fights []raid.Fight
	errors.PanicOnError(s.mapper.Order("timestamp desc").Find(&fights).Error)

	return fights
}
