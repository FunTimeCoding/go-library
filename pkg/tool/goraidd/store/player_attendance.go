package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/raid"
	"time"
)

func (s *Store) PlayerAttendance(since time.Time) []AttendanceRow {
	var available int64
	errors.PanicOnError(
		s.mapper.Model(&raid.Fight{}).
			Where("enriched = ? AND timestamp >= ?", true, since).
			Count(&available).Error,
	)
	var rows []AttendanceRow
	errors.PanicOnError(
		s.mapper.Model(&raid.PlayerFightStat{}).
			Select(
				"account",
				"string_agg(DISTINCT name, ', ' ORDER BY name) as characters",
				"count(*) as fights",
			).
			Joins("JOIN fights ON fights.filename = player_fight_stats.filename").
			Where("fights.timestamp >= ?", since).
			Group("account").
			Order("fights DESC").
			Find(&rows).Error,
	)

	for i := range rows {
		rows[i].Available = int(available)
	}

	return rows
}
