package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"time"
)

func (s *Store) MustPrune(cutoff time.Time) int {
	result, e := s.Prune(cutoff)
	errors.PanicOnError(e)

	return result
}
