package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/alias"
)

func (s *Store) GetAlias(sessionIdentifier string) string {
	var a alias.Alias
	result := s.database.Where(
		"session_identifier = ?",
		sessionIdentifier,
	).Limit(1).Find(&a)
	errors.PanicOnError(result.Error)

	if result.RowsAffected == 0 {
		return ""
	}

	return a.Name
}
