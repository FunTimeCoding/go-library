package elite_parser

import "github.com/funtimecoding/go-library/pkg/raid/elite"

func extractDefensive(player elite.Player) Defensive {
	if len(player.Defenses) == 0 {
		return Defensive{}
	}

	d := player.Defenses[0]

	return Defensive{
		DamageTaken:      d.DamageTaken,
		DownCount:        d.DownCount,
		DeadCount:        d.DeadCount,
		DodgeCount:       d.DodgeCount,
		EvadedCount:      d.EvadedCount,
		BlockedCount:     d.BlockedCount,
		InvulnedCount:    d.InvulnedCount,
		InterruptedCount: d.InterruptedCount,
	}
}
