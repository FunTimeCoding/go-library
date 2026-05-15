package mock_client

func (c *Client) Equip(key string) error {
	c.gear.Equipped["weapon"] = key

	return nil
}
