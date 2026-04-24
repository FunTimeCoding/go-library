package firefox

type request struct {
	Method     string `json:"method"`
	Parameters any    `json:"params"`
	Identifier int    `json:"id"`
}
