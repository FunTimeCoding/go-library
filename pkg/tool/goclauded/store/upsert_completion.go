package store

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/completion"

func (s *Store) UpsertCompletion(
	sessionIdentifier string,
	name string,
	kind string,
	topic string,
	message string,
) error {
	var existing completion.Completion
	result := s.database.Where(
		"session_identifier = ? AND topic = ?",
		sessionIdentifier,
		topic,
	).Limit(1).Find(&existing)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected > 0 {
		return s.database.Model(&existing).Updates(
			map[string]any{
				"name":    name,
				"kind":    kind,
				"summary": message,
			},
		).Error
	}

	return s.database.Create(
		completion.New(sessionIdentifier, name, kind, topic, message),
	).Error
}
