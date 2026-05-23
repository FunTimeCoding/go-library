package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/label"
)

func (s *Store) SetLabel(
	sessionIdentifier string,
	key string,
	value string,
) string {
	var existing label.Label
	result := s.database.Where(
		"session_identifier = ? AND key = ?",
		sessionIdentifier,
		key,
	).Limit(1).Find(&existing)
	errors.PanicOnError(result.Error)

	if result.RowsAffected > 0 {
		old := existing.Value

		if old == value {
			return old
		}

		errors.PanicOnError(
			s.database.Model(&existing).
				Update(constant.Value, value).Error,
		)

		return old
	}

	errors.PanicOnError(
		s.database.Create(label.New(sessionIdentifier, key, value)).Error,
	)

	return ""
}
