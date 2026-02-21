package connection

type subscribeCommand struct {
	Identifier uint64 `json:"id"`
	Type       string `json:"type"`
	Event      string `json:"event_type,omitempty"`
}

func (c *subscribeCommand) setIdentifier(i uint64) {
	c.Identifier = i
}
