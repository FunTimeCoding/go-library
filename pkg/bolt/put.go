package bolt

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"go.etcd.io/bbolt"
)

func (c *Client) Put(
	b *bbolt.Bucket,
	k string,
	v string,
) {
	errors.PanicOnError(b.Put([]byte(k), []byte(v)))
}
