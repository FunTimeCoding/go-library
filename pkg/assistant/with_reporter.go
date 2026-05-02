package assistant

import "github.com/funtimecoding/go-library/pkg/face"

func WithReporter(r face.Reporter) Option {
	return func(c *Client) {
		c.reporter = r
	}
}
