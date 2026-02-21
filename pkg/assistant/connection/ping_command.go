package connection

type pingCommand struct {
	Identifier uint64 `json:"id"`
	Type       string `json:"type"`
}

func (c *pingCommand) setIdentifier(i uint64) {
	c.Identifier = i
}
