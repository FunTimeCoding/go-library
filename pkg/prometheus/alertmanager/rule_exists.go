package alertmanager

func (c *Client) RuleExists(name string) bool {
	return c.Rules().Find(name) != nil
}
