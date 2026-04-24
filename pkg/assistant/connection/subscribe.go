package connection

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Connection) Subscribe(
	event string,
	s Subscriber,
) {
	_, e := c.send(&subscribeCommand{Type: Subscribe, Event: event}, s)
	errors.PanicOnError(e)
}
