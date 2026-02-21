package connection

func (c *Connection) Ping() error {
	_, e := c.sendCommand(&pingCommand{Type: Ping}, nil)

	return e
}
