package connection

func (c *Connection) Ping() error {
	_, e := c.send(&pingCommand{Type: Ping}, nil)

	return e
}
