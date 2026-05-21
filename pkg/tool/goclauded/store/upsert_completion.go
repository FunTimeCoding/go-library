package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/completion"
)

func (s *Store) UpsertCompletion(
	sessionIdentifier string,
	name string,
	kind string,
	topic string,
	message string,
) {
	var existing completion.Completion
	result := s.database.Where(
		"session_identifier = ? AND topic = ?",
		sessionIdentifier,
		topic,
	).Limit(1).Find(&existing)
	errors.PanicOnError(result.Error)

	if result.RowsAffected > 0 {
		errors.PanicOnError(
			s.database.Model(&existing).
				Updates(
					map[string]any{
						"name":    name,
						"kind":    kind,
						"summary": message,
					},
				).Error,
		)

		return
	}

	errors.PanicOnError(
		s.database.Create(
			completion.New(sessionIdentifier, name, kind, topic, message),
		).Error,
	)
}
