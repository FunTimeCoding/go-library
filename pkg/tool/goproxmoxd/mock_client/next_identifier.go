package mock_client

func (c *Client) NextIdentifier() (int, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	result := c.nextID
	c.nextID++

	return result, nil
}
