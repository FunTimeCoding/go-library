package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/store/result"
)

func (s *Store) MustGetDocument(reference string) *result.Document {
	result, e := s.GetDocument(reference)
	errors.PanicOnError(e)

	return result
}
