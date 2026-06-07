package store

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/event"

func (s *Store) UpsertEvent(
	sessionIdentifier string,
	kind string,
	name string,
	target string,
	body string,
) error {
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

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected > 0 {
		return s.database.Model(&existing).Updates(
			map[string]any{
				"name":   name,
				"target": target,
				"body":   body,
			},
		).Error
	}

	return s.LogEvent(sessionIdentifier, kind, name, target, body)
}
