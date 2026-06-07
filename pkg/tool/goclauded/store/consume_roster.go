package store

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"

func (s *Store) ConsumeRoster(name string) (bool, error) {
	var i session.Session
	result := s.database.Where("callsign = ?", name).Limit(1).Find(&i)

	if result.Error != nil {
		return false, result.Error
	}

	if result.RowsAffected == 0 || !i.NeedsRoster {
		return false, nil
	}

	if e := s.database.Model(&i).UpdateColumn(
		"needs_roster",
		false,
	).Error; e != nil {
		return false, e
	}

	return true, nil
}
