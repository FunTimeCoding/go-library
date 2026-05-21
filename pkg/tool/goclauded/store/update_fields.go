package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) UpdateFields(
	identifier string,
	updates map[string]any,
) {
	errors.PanicOnError(
		s.database.Model(session.Stub()).
			Where("identifier = ?", identifier).
			Updates(updates).Error,
	)
}
