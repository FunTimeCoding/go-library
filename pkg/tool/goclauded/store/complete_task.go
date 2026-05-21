package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) CompleteTask(name string) string {
	var i session.Session
	result := s.database.Where("callsign = ?", name).Limit(1).Find(&i)
	errors.PanicOnError(result.Error)

	if result.RowsAffected == 0 {
		return ""
	}

	topic := i.Topic
	errors.PanicOnError(
		s.database.Model(session.Stub()).
			Where("callsign = ?", name).
			Updates(
				map[string]any{
					"topic": "",
					"files": "",
				},
			).Error,
	)

	return topic
}
