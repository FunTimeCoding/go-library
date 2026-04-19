package elite

type Fight struct {
	TimeStartStd string   `json:"timeStartStd"`
	DurationMS   int      `json:"durationMS"`
	MapID        int      `json:"mapID"`
	FightName    string   `json:"fightName"`
	Success      bool     `json:"success"`
	Targets      []Target `json:"targets"`
	Players      []Player `json:"players"`
}
