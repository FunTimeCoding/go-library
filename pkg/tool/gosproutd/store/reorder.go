package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/store/seed"
)

func (s *Store) Reorder(identifiers []uint) {
	for i, v := range identifiers {
		errors.PanicOnError(
			s.mapper.Model(&seed.Seed{Identifier: v}).
				Update("position", i+1).Error,
		)
	}
}
