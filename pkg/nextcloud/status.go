package nextcloud

func (c *Client) Status() {
	c.basic.Propfind()
}
