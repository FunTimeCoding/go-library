package store

import "github.com/funtimecoding/go-library/pkg/errors"

func (s *Store) ContentCount() int {
	var count int
	errors.PanicOnError(
		s.database.QueryRow("SELECT COUNT(*) FROM content").Scan(&count),
	)

	return count
}
