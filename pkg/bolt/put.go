package bolt

import "go.etcd.io/bbolt"

func (c *Client) Put(
	b *bbolt.Bucket,
	k string,
	v string,
) error {
	return b.Put([]byte(k), []byte(v))
}
