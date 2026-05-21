package store

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"

func (s *Store) isListening(name string) bool {
	var i session.Session
	result := s.database.Where("callsign = ?", name).Limit(1).Find(&i)

	if result.Error != nil || result.RowsAffected == 0 {
		return false
	}

	return i.Listening
}
