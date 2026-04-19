package store

import "github.com/funtimecoding/go-library/pkg/errors"

func (s *Store) PlayerDetail(account string) []ProfessionRow {
	var rows []ProfessionRow
	errors.PanicOnError(
		s.mapper.
			Table("player_fight_stats").
			Select(
				"profession",
				"count(*) as fights",
				"sum(damage) as damage",
				"sum(active_time_ms) as active_time_ms",
				"sum(dead_count) as dead_count",
				"sum(boon_strips) as boon_strips",
			).
			Where("account = ?", account).
			Group("profession").
			Order("fights DESC").
			Find(&rows).Error,
	)

	return rows
}
