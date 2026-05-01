package bolt

import "go.etcd.io/bbolt"

func (c *Client) PutBytes(
	b *bbolt.Bucket,
	k string,
	v []byte,
) error {
	return b.Put([]byte(k), v)
}
