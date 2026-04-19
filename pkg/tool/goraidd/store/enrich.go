package store

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/raid"
	"github.com/funtimecoding/go-library/pkg/raid/elite"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (s *Store) enrichFile(path string) {
	name := filepath.Base(path)

	if !strings.HasSuffix(name, "_detailed_wvw_kill.json") {
		return
	}

	data, e := os.ReadFile(path)

	if e != nil {
		slog.Error("enrich read failed", "path", path, "error", e)

		return
	}

	var fight elite.Fight

	if e := json.Unmarshal(data, &fight); e != nil {
		slog.Error("enrich parse failed", "path", path, "error", e)

		return
	}

	timestamp, e := time.Parse(
		"2006-01-02 15:04:05 -07:00",
		fight.TimeStartStd,
	)

	if e != nil {
		slog.Error(
			"enrich timestamp parse failed",
			"value",
			fight.TimeStartStd,
			"error",
			e,
		)

		return
	}

	alliedTeamID := 0

	if len(fight.Players) > 0 {
		alliedTeamID = fight.Players[0].TeamID
	}

	enemyCount := 0
	enemyTeams := map[int]int{}

	for _, target := range fight.Targets {
		if target.TeamID == 0 {
			continue
		}

		enemyCount++
		enemyTeams[target.TeamID]++
	}

	enemyTeamsJSON, marshalError := json.Marshal(enemyTeams)
	errors.PanicOnError(marshalError)
	enemyTeamsString := string(enemyTeamsJSON)
	zevtcBase := strings.TrimSuffix(name, "_detailed_wvw_kill.json")
	slog.Info("enriching fight", "file", name, "zevtc", zevtcBase)
	result := s.mapper.Model(&raid.Fight{}).
		Where("filename LIKE ?", "%"+zevtcBase+"%").
		Updates(
			map[string]any{
				"timestamp":      timestamp,
				"duration_ms":    fight.DurationMS,
				"map_id":         fight.MapID,
				"map_name":       fight.FightName,
				"allied_count":   len(fight.Players),
				"allied_team_id": alliedTeamID,
				"enemy_count":    enemyCount,
				"enemy_teams":    enemyTeamsString,
				"success":        fight.Success,
				"enriched":       true,
			},
		)

	if result.RowsAffected == 0 {
		slog.Warn("enrich: no matching fight row", "zevtc", zevtcBase)
	}
}
