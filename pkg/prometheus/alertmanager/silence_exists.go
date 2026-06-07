package alertmanager

func (c *Client) SilenceExists(name string) (bool, error) {
	v, e := c.Silences(false)

	if e != nil {
		return false, e
	}

	for _, s := range v {
		if s.Rule == name {
			return true, nil
		}
	}

	return false, nil
}
