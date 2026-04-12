package firefox

import (
	"context"
	"fmt"
	"github.com/coder/websocket/wsjson"
	"time"
)

func (c *Client) send(
	method string,
	parameters any,
) (response, error) {
	c.mutex.Lock()
	connection := c.connection
	c.mutex.Unlock()

	if connection == nil {
		return response{}, fmt.Errorf("extension not connected")
	}

	identifier := int(c.identifier.Add(1))
	channel := make(chan response, 1)
	c.mutex.Lock()
	c.pending[identifier] = channel
	c.mutex.Unlock()

	defer func() {
		c.mutex.Lock()
		delete(c.pending, identifier)
		c.mutex.Unlock()
	}()
	e := wsjson.Write(
		context.Background(),
		connection,
		request{
			Method:     method,
			Parameters: parameters,
			Identifier: identifier,
		},
	)

	if e != nil {
		return response{}, fmt.Errorf("%s: %w", method, e)
	}

	select {
	case r := <-channel:
		if r.Error != "" {
			return response{}, fmt.Errorf("%s: %s", method, r.Error)
		}

		return r, nil
	case <-time.After(10 * time.Second):
		return response{}, fmt.Errorf(
			"%s: timeout waiting for response",
			method,
		)
	}
}
