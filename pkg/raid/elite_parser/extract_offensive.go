package elite_parser

import "github.com/funtimecoding/go-library/pkg/raid/elite"

func extractOffensive(player elite.Player) Offensive {
	var result Offensive

	for _, targets := range player.DPSTargets {
		for _, target := range targets {
			result.Damage += target.Damage
		}
	}

	for _, targets := range player.StatsTargets {
		for _, target := range targets {
			result.Kills += target.Killed
			result.Downs += target.Downed
		}
	}

	return result
}
