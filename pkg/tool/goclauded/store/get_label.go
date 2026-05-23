package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/label"
)

func (s *Store) GetLabel(
	sessionIdentifier string,
	key string,
) string {
	var existing label.Label
	result := s.database.Where(
		"session_identifier = ? AND key = ?",
		sessionIdentifier,
		key,
	).Limit(1).Find(&existing)
	errors.PanicOnError(result.Error)

	if result.RowsAffected == 0 {
		return ""
	}

	return existing.Value
}
