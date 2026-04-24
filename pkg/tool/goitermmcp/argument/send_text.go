package argument

type SendText struct {
	SessionIdentifier string `json:"session_id"`
	Text              string `json:"text"`
	SendEnter         bool   `json:"send_enter"`
}
