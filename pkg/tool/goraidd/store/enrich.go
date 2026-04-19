package store

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/raid"
	"github.com/funtimecoding/go-library/pkg/raid/elite"
	"github.com/funtimecoding/go-library/pkg/raid/elite_parser"
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
	var fightRow raid.Fight
	lookup := s.mapper.
		Where("filename LIKE ?", "%"+zevtcBase+"%").
		First(&fightRow)

	if lookup.Error != nil {
		slog.Warn("enrich: no matching fight row", "zevtc", zevtcBase)

		return
	}

	s.mapper.Model(&fightRow).Updates(
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

	if !elite_parser.IsValidFight(&fight) {
		return
	}

	s.mapper.Where("filename = ?", fightRow.Filename).
		Delete(&raid.PlayerFightStat{})
	stats := elite_parser.Extract(&fight)

	for _, stat := range stats {
		row := raid.PlayerFightStat{
			Filename:          fightRow.Filename,
			Account:           stat.Identity.Account,
			Name:              stat.Identity.Name,
			Profession:        stat.Identity.Profession,
			GroupNumber:       stat.Identity.Group,
			Commander:         stat.Identity.Commander,
			ActiveTimeMS:      stat.Identity.ActiveTimeMS,
			Damage:            stat.Offensive.Damage,
			DamageTaken:       stat.Defensive.DamageTaken,
			DownCount:         stat.Defensive.DownCount,
			DeadCount:         stat.Defensive.DeadCount,
			Kills:             stat.Offensive.Kills,
			Downs:             stat.Offensive.Downs,
			BoonStrips:        stat.Support.BoonStrips,
			ConditionCleanses: stat.Support.ConditionCleanses,
			Healing:           stat.Support.Healing,
			Barrier:           stat.Support.Barrier,
			DodgeCount:        stat.Defensive.DodgeCount,
			EvadedCount:       stat.Defensive.EvadedCount,
			BlockedCount:      stat.Defensive.BlockedCount,
			InvulnedCount:     stat.Defensive.InvulnedCount,
			StabilityUptime:   stat.Boons.Stability,
			MightUptime:       stat.Boons.Might,
			FuryUptime:        stat.Boons.Fury,
			QuicknessUptime:   stat.Boons.Quickness,
			ProtectionUptime:  stat.Boons.Protection,
			ResistanceUptime:  stat.Boons.Resistance,
			DistToCom:         stat.Identity.DistToCom,
		}
		errors.PanicOnError(s.mapper.Create(&row).Error)
	}

	slog.Info(
		"enriched player stats",
		"file",
		name,
		"players",
		len(stats),
	)
}
