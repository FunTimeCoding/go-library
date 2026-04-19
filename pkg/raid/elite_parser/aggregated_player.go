package elite_parser

type AggregatedPlayer struct {
	Account           string
	Name              string
	Profession        string
	Fights            int
	Offensive         Offensive
	Defensive         Defensive
	Support           Support
	Boons             Boons
	TotalActiveTimeMS int
}
