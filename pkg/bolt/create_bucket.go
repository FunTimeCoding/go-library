package bolt

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"go.etcd.io/bbolt"
)

func (c *Client) CreateBucket(
	t *bbolt.Tx,
	name string,
) *bbolt.Bucket {
	result, e := t.CreateBucketIfNotExists([]byte(name))
	errors.PanicOnError(e)

	return result
}
