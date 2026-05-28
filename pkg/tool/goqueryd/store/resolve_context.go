package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func (s *Store) resolveContext(
	collection string,
	filePath string,
) string {
	normalized := join.Empty(separator.Slash, filePath)
	prefixes := pathPrefixes(normalized)
	rows, e := s.database.Query(
		`SELECT path_prefix, description FROM context
		WHERE collection = ? ORDER BY length(path_prefix)`,
		collection,
	)
	errors.PanicOnError(e)
	defer errors.PanicClose(rows)
	var descriptions []string

	for rows.Next() {
		var prefix, description string
		errors.PanicOnError(rows.Scan(&prefix, &description))

		if prefixes[prefix] {
			descriptions = append(descriptions, description)
		}
	}

	errors.PanicOnError(rows.Err())

	return strings.Join(descriptions, "\n\n")
}
