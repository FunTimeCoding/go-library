package confluence

type commentPayload struct {
	Type      string           `json:"type"`
	Container commentContainer `json:"container"`
	Body      commentBody      `json:"body"`
}
