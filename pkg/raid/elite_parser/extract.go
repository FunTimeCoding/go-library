package elite_parser

import "github.com/funtimecoding/go-library/pkg/raid/elite"

func Extract(fight *elite.Fight) []PlayerFightStat {
	var result []PlayerFightStat

	for _, player := range fight.Players {
		if player.NotInSquad {
			continue
		}

		stat := PlayerFightStat{
			Identity:  extractIdentity(player),
			Offensive: extractOffensive(player),
			Defensive: extractDefensive(player),
			Support:   extractSupport(player),
			Boons:     extractBoons(player),
		}
		result = append(result, stat)
	}

	return result
}
