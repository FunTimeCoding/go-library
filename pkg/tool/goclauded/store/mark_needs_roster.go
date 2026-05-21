package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) MarkNeedsRoster(identifier string) {
	errors.PanicOnError(
		s.database.Model(session.Stub()).
			Where("identifier = ?", identifier).
			UpdateColumn("needs_roster", true).Error,
	)
}
