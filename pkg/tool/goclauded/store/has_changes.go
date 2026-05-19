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
) bool {
	var e session.Session
	s.database.Where("name = ?", name).Limit(1).Find(&e)

	if e.NeedsRoster {
		s.database.Model(&e).UpdateColumn("needs_roster", false)

		return true
	}

	var rosterCount int64
	s.database.Model(session.New()).
		Where("name != ? AND updated_at > ?", name, lastSeen).
		Count(&rosterCount)

	if rosterCount > 0 {
		return true
	}

	var messageCount int64
	s.database.Model(message.Stub()).
		Where("(to_name = ? OR to_name = '') AND read = ?", name, false).
		Count(&messageCount)

	if messageCount > 0 {
		return true
	}

	var activityCount int64
	s.database.Model(completion.Stub()).
		Where("created_at > ?", lastSeen).
		Count(&activityCount)

	return activityCount > 0
}
