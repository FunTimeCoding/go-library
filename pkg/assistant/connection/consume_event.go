package connection

import "github.com/funtimecoding/go-library/pkg/assistant/message"

func (c *Connection) consumeEvent(m *message.Message) {
	c.RLock()
	f, ok := c.subscribers[m.Identifier]
	c.RUnlock()

	if !ok {
		return
	}

	go f(m)
}
