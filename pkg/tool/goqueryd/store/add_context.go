package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"strings"
)

func (s *Store) AddContext(
	collection string,
	pathPrefix string,
	description string,
) {
	if !strings.HasPrefix(pathPrefix, "/") {
		pathPrefix = join.Empty("/", pathPrefix)
	}

	if !strings.HasSuffix(pathPrefix, "/") {
		pathPrefix = join.Empty(pathPrefix, "/")
	}

	_, e := s.database.Exec(
		`INSERT INTO context (collection, path_prefix, description)
		VALUES (?, ?, ?)
		ON CONFLICT(collection, path_prefix) DO UPDATE SET
			description = excluded.description`,
		collection,
		pathPrefix,
		description,
	)
	errors.PanicOnError(e)
}
