package netbox

func WithVerbose(v bool) OptionFunc {
	return func(c *Client) {
		c.verbose = v
	}
}
