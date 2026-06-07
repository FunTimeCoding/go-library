package store

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/completion"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/message"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
	"time"
)

func (s *Store) HasChanges(
	name string,
	lastSeen time.Time,
) (bool, error) {
	var rosterCount int64

	if e := s.database.Model(session.Stub()).Where(
		"name != ? AND last_seen > ?",
		name,
		lastSeen,
	).Count(&rosterCount).Error; e != nil {
		return false, e
	}

	if rosterCount > 0 {
		return true, nil
	}

	var messageCount int64

	if e := s.database.Model(message.Stub()).Where(
		"(to_name = ? OR to_name = '') AND read = ?",
		name,
		false,
	).Count(&messageCount).Error; e != nil {
		return false, e
	}

	if messageCount > 0 {
		return true, nil
	}

	var activityCount int64

	if e := s.database.Model(completion.Stub()).Where(
		"created_at > ?",
		lastSeen,
	).Count(&activityCount).Error; e != nil {
		return false, e
	}

	return activityCount > 0, nil
}
