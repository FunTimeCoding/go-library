package connection

import "context"

func New(
	host string,
	token string,
) *Connection {
	x, cancel := context.WithCancel(context.Background())

	return &Connection{
		host:        host,
		token:       token,
		subscribers: make(map[uint64]Subscriber),
		context:     x,
		cancel:      cancel,
	}
}
