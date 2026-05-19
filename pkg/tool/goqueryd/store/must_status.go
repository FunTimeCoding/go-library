package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/store/status"
)

func (s *Store) MustStatus() *status.Status {
	r, e := s.Status()
	errors.PanicOnError(e)

	return r
}
