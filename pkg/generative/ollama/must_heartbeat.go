package ollama

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustHeartbeat() {
	errors.PanicOnError(c.Heartbeat())
}
