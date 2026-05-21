package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/event"
)

func (s *Store) UpsertEvent(
	sessionIdentifier string,
	kind string,
	name string,
	target string,
	body string,
) {
	var existing event.Event
	query := s.database.Where(
		"session_identifier = ? AND kind = ?",
		sessionIdentifier,
		kind,
	)

	if target != "" {
		query = query.Where("target = ?", target)
	}

	result := query.Limit(1).Find(&existing)
	errors.PanicOnError(result.Error)

	if result.RowsAffected > 0 {
		errors.PanicOnError(
			s.database.Model(&existing).
				Updates(
					map[string]any{
						"name":   name,
						"target": target,
						"body":   body,
					},
				).Error,
		)

		return
	}

	s.LogEvent(sessionIdentifier, kind, name, target, body)
}
