package exceptions

type Exception struct {
	Name   string `json:"Name"`
	Guild  string `json:"Guild"`
	Reason string `json:"Reason"`
}
