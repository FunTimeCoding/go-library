package bolt

import "go.etcd.io/bbolt"

func (c *Client) Get(
	b *bbolt.Bucket,
	k string,
) string {
	return string(b.Get([]byte(k)))
}
