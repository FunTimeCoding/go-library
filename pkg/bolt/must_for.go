package bolt

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"go.etcd.io/bbolt"
)

func MustFor(
	b *bbolt.Bucket,
	f func(
		k string,
		v []byte,
	),
) {
	errors.PanicOnError(For(b, f))
}
