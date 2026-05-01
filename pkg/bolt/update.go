package bolt

import "go.etcd.io/bbolt"

func (c *Client) Update(f func(t *bbolt.Tx) error) error {
	return c.client.Update(f)
}
