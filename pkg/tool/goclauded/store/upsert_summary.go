package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/summary"
)

func (s *Store) UpsertSummary(
	sessionIdentifier string,
	name string,
	body string,
) {
	var existing summary.Summary
	result := s.database.Where(
		"session_identifier = ?",
		sessionIdentifier,
	).Limit(1).Find(&existing)
	errors.PanicOnError(result.Error)

	if result.RowsAffected > 0 {
		errors.PanicOnError(
			s.database.Model(&existing).
				Updates(
					map[string]any{
						"name": name,
						"body": body,
					},
				).Error,
		)
		s.notify()

		return
	}

	errors.PanicOnError(
		s.database.Create(summary.New(sessionIdentifier, name, body)).Error,
	)
	s.notify()
}
