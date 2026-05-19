package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) ConsumeReannounce(name string) bool {
	var i session.Session
	result := s.database.Where("name = ?", name).Limit(1).Find(&i)

	if result.Error != nil || result.RowsAffected == 0 || !i.NeedsReannounce {
		return false
	}

	errors.PanicOnError(
		s.database.Model(session.New()).
			Where("name = ?", name).
			Update("needs_reannounce", false).Error,
	)

	return true
}
