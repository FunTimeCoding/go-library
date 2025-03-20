package alertmanager

func (c *Client) RuleExists(name string) bool {
	return c.Rule(name) != nil
}
