package store

import "time"

func (s *Store) Recent(limit int) ([]Run, error) {
	var result []Run

	return result, s.mapper.
		Table(s.tableName).
		Where("created_at > ?", time.Now().Add(-RetentionAge)).
		Order("created_at desc").
		Limit(limit).
		Find(&result).Error
}
