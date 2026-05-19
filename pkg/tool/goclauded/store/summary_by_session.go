package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/summary"
)

func (s *Store) SummaryBySession(sessionIdentifier string) string {
	var i summary.Summary
	result := s.database.Where(
		"session_identifier = ?",
		sessionIdentifier,
	).Limit(1).Find(&i)
	errors.PanicOnError(result.Error)

	if result.RowsAffected == 0 {
		return ""
	}

	return i.Body
}
