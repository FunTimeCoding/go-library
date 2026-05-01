package bolt

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"go.etcd.io/bbolt"
)

func (c *Client) MustPut(
	b *bbolt.Bucket,
	k string,
	v string,
) {
	errors.PanicOnError(c.Put(b, k, v))
}
