package alertmanager

func (c *Client) SimpleSilence(alert string) (string, error) {
	return c.SetSilence(alert, "", 0)
}
