package alertmanager

func (c *Client) SimpleSilence(alert string) string {
	return c.SetSilence(alert, "", 0)
}
