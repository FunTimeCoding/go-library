package ollama

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) Heartbeat() {
	errors.PanicOnError(c.client.Heartbeat(c.context))
}
