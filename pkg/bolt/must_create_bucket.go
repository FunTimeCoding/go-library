package bolt

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"go.etcd.io/bbolt"
)

func (c *Client) MustCreateBucket(
	t *bbolt.Tx,
	name string,
) *bbolt.Bucket {
	result, e := c.CreateBucket(t, name)
	errors.PanicOnError(e)

	return result
}
