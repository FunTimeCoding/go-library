package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/store/result"
)

func (s *Store) MustListDocuments(collection string) []result.DocumentEntry {
	result, e := s.ListDocuments(collection)
	errors.PanicOnError(e)

	return result
}
