package bolt

import "go.etcd.io/bbolt"

func (c *Client) View(f func(t *bbolt.Tx) error) error {
	return c.client.View(f)
}
