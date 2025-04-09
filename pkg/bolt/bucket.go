package bolt

import "go.etcd.io/bbolt"

func (c *Client) Bucket(
	t *bbolt.Tx,
	name string,
) *bbolt.Bucket {
	return t.Bucket([]byte(name))
}
