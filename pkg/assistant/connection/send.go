package connection

import "fmt"

func (c *Connection) sendCommand(
	o command,
	s Subscriber,
) (uint64, error) {
	c.Lock()
	defer c.Unlock()

	if c.connection == nil {
		return 0, fmt.Errorf("connection not open")
	}

	c.lastIdentifier++
	o.setIdentifier(c.lastIdentifier)

	if s != nil {
		c.subscribers[c.lastIdentifier] = s
	}

	if e := c.connection.WriteJSON(o); e != nil {
		delete(c.subscribers, c.lastIdentifier)

		return 0, fmt.Errorf("send fail: %w", e)
	}

	return c.lastIdentifier, nil
}
