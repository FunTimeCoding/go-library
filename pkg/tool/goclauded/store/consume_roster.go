package store

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"

func (s *Store) ConsumeRoster(name string) bool {
	var i session.Session
	s.database.Where("callsign = ?", name).Limit(1).Find(&i)

	if !i.NeedsRoster {
		return false
	}

	s.database.Model(&i).UpdateColumn("needs_roster", false)

	return true
}
