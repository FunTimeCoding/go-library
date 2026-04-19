package elite_parser

import "github.com/funtimecoding/go-library/pkg/raid/elite"

func extractSupport(player elite.Player) Support {
	var result Support

	if len(player.Support) > 0 {
		result.BoonStrips = player.Support[0].BoonStrips
		result.ConditionCleanses = player.Support[0].CondiCleanse
		result.Resurrects = player.Support[0].Resurrects
		result.ResurrectTime = player.Support[0].ResurrectTime
		result.StunBreak = player.Support[0].StunBreak
	}

	if player.ExtHealing != nil && len(player.ExtHealing.OutgoingHealing) > 0 {
		result.Healing = player.ExtHealing.OutgoingHealing[0].Healing
	}

	if player.ExtBarrier != nil && len(player.ExtBarrier.OutgoingBarrier) > 0 {
		result.Barrier = player.ExtBarrier.OutgoingBarrier[0].Barrier
	}

	return result
}
