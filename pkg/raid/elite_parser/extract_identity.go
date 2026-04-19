package elite_parser

import "github.com/funtimecoding/go-library/pkg/raid/elite"

func extractIdentity(player elite.Player) Identity {
	result := Identity{
		Account:    player.Account,
		Name:       player.Name,
		Profession: player.Profession,
		Group:      player.Group,
		Commander:  player.HasCommanderTag,
	}

	if len(player.ActiveTimes) > 0 {
		result.ActiveTimeMS = player.ActiveTimes[0]
	}

	if len(player.StatsAll) > 0 {
		result.DistToCom = player.StatsAll[0].DistToCom
	}

	return result
}
