package bolt

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"go.etcd.io/bbolt"
)

func (c *Client) MustUpdate(f func(t *bbolt.Tx) error) {
	errors.PanicOnError(c.Update(f))
}
