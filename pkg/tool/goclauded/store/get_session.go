package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) GetSession(identifier string) *session.Session {
	var i session.Session
	result := s.database.Where("identifier = ?", identifier).Limit(1).Find(&i)
	errors.PanicOnError(result.Error)

	if result.RowsAffected == 0 {
		return nil
	}

	return &i
}
