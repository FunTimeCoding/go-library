package store

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/raid"
	"github.com/funtimecoding/go-library/pkg/raid/elite"
	"github.com/funtimecoding/go-library/pkg/raid/elite_parser"
	"github.com/funtimecoding/go-library/pkg/strings/join"
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

	b, e := os.ReadFile(path)

	if e != nil {
		slog.Error("enrich read failed", "path", path, "error", e)

		return
	}

	var fight elite.Fight

	if e := json.Unmarshal(b, &fight); e != nil {
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
		Where("filename LIKE ?", join.Empty("%", zevtcBase, "%")).
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
		Delete(raid.NewPlayerFightStat())
	stats := elite_parser.Extract(&fight)

	for _, stat := range stats {
		row := raid.NewPlayerFightStat()
		row.Filename = fightRow.Filename
		row.Account = stat.Identity.Account
		row.Name = stat.Identity.Name
		row.Profession = stat.Identity.Profession
		row.GroupNumber = stat.Identity.Group
		row.Commander = stat.Identity.Commander
		row.ActiveTimeMS = stat.Identity.ActiveTimeMS
		row.Damage = stat.Offensive.Damage
		row.DamageTaken = stat.Defensive.DamageTaken
		row.DownCount = stat.Defensive.DownCount
		row.DeadCount = stat.Defensive.DeadCount
		row.Kills = stat.Offensive.Kills
		row.Downs = stat.Offensive.Downs
		row.BoonStrips = stat.Support.BoonStrips
		row.ConditionCleanses = stat.Support.ConditionCleanses
		row.Healing = stat.Support.Healing
		row.Barrier = stat.Support.Barrier
		row.DodgeCount = stat.Defensive.DodgeCount
		row.EvadedCount = stat.Defensive.EvadedCount
		row.BlockedCount = stat.Defensive.BlockedCount
		row.InvulnedCount = stat.Defensive.InvulnedCount
		row.StabilityUptime = stat.Boons.Stability
		row.MightUptime = stat.Boons.Might
		row.FuryUptime = stat.Boons.Fury
		row.QuicknessUptime = stat.Boons.Quickness
		row.ProtectionUptime = stat.Boons.Protection
		row.ResistanceUptime = stat.Boons.Resistance
		row.DistToCom = stat.Identity.DistToCom
		errors.PanicOnError(s.mapper.Create(row).Error)
	}

	slog.Info(
		"enriched player stats",
		"file",
		name,
		"players",
		len(stats),
	)
}
