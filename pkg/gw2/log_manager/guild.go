package log_manager

type Guild struct {
	Id     string `json:"Id"`
	Name   string `json:"Name"`
	Tag    string `json:"Tag"`
	Emblem struct {
		Background Face  `json:"Background"`
		Foreground Face  `json:"Foreground"`
		Flags      []any `json:"Flags"`
	} `json:"Emblem"`
}
