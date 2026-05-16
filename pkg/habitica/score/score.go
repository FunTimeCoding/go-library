package score

type Score struct {
	HP    float64 `json:"hp"`
	MP    float64 `json:"mp"`
	XP    float64 `json:"exp"`
	GP    float64 `json:"gp"`
	Level int     `json:"lvl"`
}
