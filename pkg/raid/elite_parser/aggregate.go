package elite_parser

func Aggregate(stats []PlayerFightStat) map[string]*AggregatedPlayer {
	players := map[string]*AggregatedPlayer{}

	for _, stat := range stats {
		key := stat.Identity.Account + "|" + stat.Identity.Profession
		p, exists := players[key]

		if !exists {
			p = &AggregatedPlayer{
				Account:    stat.Identity.Account,
				Name:       stat.Identity.Name,
				Profession: stat.Identity.Profession,
			}
			players[key] = p
		}

		p.Fights++
		p.TotalActiveTimeMS += stat.Identity.ActiveTimeMS
		p.Offensive.Damage += stat.Offensive.Damage
		p.Offensive.Kills += stat.Offensive.Kills
		p.Offensive.Downs += stat.Offensive.Downs
		p.Defensive.DamageTaken += stat.Defensive.DamageTaken
		p.Defensive.DownCount += stat.Defensive.DownCount
		p.Defensive.DeadCount += stat.Defensive.DeadCount
		p.Defensive.DodgeCount += stat.Defensive.DodgeCount
		p.Defensive.EvadedCount += stat.Defensive.EvadedCount
		p.Defensive.BlockedCount += stat.Defensive.BlockedCount
		p.Defensive.InvulnedCount += stat.Defensive.InvulnedCount
		p.Defensive.InterruptedCount += stat.Defensive.InterruptedCount
		p.Support.BoonStrips += stat.Support.BoonStrips
		p.Support.ConditionCleanses += stat.Support.ConditionCleanses
		p.Support.Healing += stat.Support.Healing
		p.Support.Barrier += stat.Support.Barrier
		p.Support.Resurrects += stat.Support.Resurrects
		p.Support.ResurrectTime += stat.Support.ResurrectTime
		p.Support.StunBreak += stat.Support.StunBreak
		p.Boons.Stability += stat.Boons.Stability
		p.Boons.Might += stat.Boons.Might
		p.Boons.Fury += stat.Boons.Fury
		p.Boons.Quickness += stat.Boons.Quickness
		p.Boons.Protection += stat.Boons.Protection
		p.Boons.Resistance += stat.Boons.Resistance
		p.Boons.Aegis += stat.Boons.Aegis
		p.Boons.Resolution += stat.Boons.Resolution
		p.Boons.Swiftness += stat.Boons.Swiftness
		p.Boons.Vigor += stat.Boons.Vigor
		p.Boons.Regeneration += stat.Boons.Regeneration
		p.Boons.Alacrity += stat.Boons.Alacrity
	}

	return players
}
