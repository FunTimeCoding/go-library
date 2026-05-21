package store

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"

func (s *Store) SessionByCallsign(c string) *session.Session {
	var i session.Session
	result := s.database.Where("callsign = ?", c).Limit(1).Find(&i)

	if result.Error != nil || result.RowsAffected == 0 {
		return nil
	}

	return &i
}
