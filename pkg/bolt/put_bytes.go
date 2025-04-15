package bolt

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"go.etcd.io/bbolt"
)

func (c *Client) PutBytes(
	b *bbolt.Bucket,
	k string,
	v []byte,
) {
	errors.PanicOnError(b.Put([]byte(k), v))
}
