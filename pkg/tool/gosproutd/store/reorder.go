package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/store/seed"
)

func (s *Store) Reorder(identifiers []uint) {
	for i, v := range identifiers {
		errors.PanicOnError(
			s.mapper.Model(seed.Stub()).
				Where("identifier = ?", v).
				Update("position", i+1).Error,
		)
	}
}
