package store

type ProfessionRow struct {
	Profession        string
	Fights            int
	Damage            int
	Healing           int
	ConditionCleanses int
	BoonStrips        int
	Downs             int
	DeadCount         int
	ActiveTimeMS      int
	DistToCom         float64
}
