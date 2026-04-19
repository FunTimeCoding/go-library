package store

type detailedFight struct {
	TimeStartStd string           `json:"timeStartStd"`
	DurationMS   int              `json:"durationMS"`
	MapID        int              `json:"mapID"`
	FightName    string           `json:"fightName"`
	Success      bool             `json:"success"`
	Targets      []detailedTarget `json:"targets"`
	Players      []detailedPlayer `json:"players"`
}
