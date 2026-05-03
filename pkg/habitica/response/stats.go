package response

type Stats struct {
	HP     float64 `json:"hp"`
	MP     float64 `json:"mp"`
	XP     float64 `json:"exp"`
	GP     float64 `json:"gp"`
	Level  int     `json:"lvl"`
	Class  string  `json:"class"`
	Points int     `json:"points"`
	Str    int     `json:"str"`
	Con    int     `json:"con"`
	Int    int     `json:"int"`
	Per    int     `json:"per"`
}
