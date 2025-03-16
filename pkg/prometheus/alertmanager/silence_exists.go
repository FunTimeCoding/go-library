package alertmanager

func (c *Client) SilenceExists(name string) bool {
	for _, s := range c.Silences(false) {
		if s.Rule == name {
			return true
		}
	}

	return false
}
