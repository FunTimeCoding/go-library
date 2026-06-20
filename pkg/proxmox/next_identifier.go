package proxmox

func (c *Client) NextIdentifier() (int, error) {
	cluster, e := c.Cluster()

	if e != nil {
		return 0, e
	}

	return cluster.NextID(c.context)
}
