package habitica

type CreateTaskBody struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Notes string `json:"notes,omitempty"`
}
