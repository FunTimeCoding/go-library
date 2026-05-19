package store

import "github.com/funtimecoding/go-library/pkg/errors"

func (s *Store) MustSearchKeyword(
	query string,
	limit int,
	collection string,
	full bool,
) []SearchResult {
	result, e := s.SearchKeyword(query, limit, collection, full)
	errors.PanicOnError(e)

	return result
}
