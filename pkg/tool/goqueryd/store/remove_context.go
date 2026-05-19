package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"strings"
)

func (s *Store) RemoveContext(
	collection string,
	pathPrefix string,
) bool {
	if !strings.HasPrefix(pathPrefix, "/") {
		pathPrefix = join.Empty("/", pathPrefix)
	}

	if !strings.HasSuffix(pathPrefix, "/") {
		pathPrefix = join.Empty(pathPrefix, "/")
	}

	result, e := s.database.Exec(
		"DELETE FROM context WHERE collection = ? AND path_prefix = ?",
		collection,
		pathPrefix,
	)
	errors.PanicOnError(e)
	affected, e := result.RowsAffected()
	errors.PanicOnError(e)

	return affected > 0
}
