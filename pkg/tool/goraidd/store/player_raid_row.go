package store

import "time"

type PlayerRaidRow struct {
	RaidID            uint
	RaidName          string
	RaidDate          time.Time
	RaidFights        int
	Profession        string
	Fights            int
	Damage            int
	Healing           int
	ConditionCleanses int
	BoonStrips        int
	Barrier           int
	Downs             int
	DeadCount         int
	ActiveTimeMS      int
	DistToCom         float64
}
