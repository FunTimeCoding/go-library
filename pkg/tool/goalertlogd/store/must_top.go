package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"time"
)

func (s *Store) MustTop(
	n int,
	start time.Time,
	end time.Time,
) []TopRecord {
	result, e := s.Top(n, start, end)
	errors.PanicOnError(e)

	return result
}
