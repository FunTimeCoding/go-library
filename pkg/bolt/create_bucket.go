package bolt

import "go.etcd.io/bbolt"

func (c *Client) CreateBucket(
	t *bbolt.Tx,
	name string,
) (*bbolt.Bucket, error) {
	return t.CreateBucketIfNotExists([]byte(name))
}
