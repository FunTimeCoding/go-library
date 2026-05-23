package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/label"
)

func (s *Store) DeleteLabel(
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

	old := existing.Value
	errors.PanicOnError(
		s.database.Delete(&existing).Error,
	)

	return old
}
