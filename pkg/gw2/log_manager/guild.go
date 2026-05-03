package log_manager

type Guild struct {
	Id     string `json:"Id"`
	Name   string `json:"Name"`
	Tag    string `json:"Tag"`
	Emblem Emblem `json:"Emblem"`
}
