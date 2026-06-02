package convert

type SlimInstance struct {
	Name   string `json:"name"`
	Host   string `json:"host"`
	Port   int    `json:"port"`
	Active bool   `json:"active"`
}
