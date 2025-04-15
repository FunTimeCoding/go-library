package bolt

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"go.etcd.io/bbolt"
)

func For(
	b *bbolt.Bucket,
	f func(
		k string,
		v []byte,
	),
) {
	errors.PanicOnError(
		b.ForEach(
			func(
				k []byte,
				v []byte,
			) error {
				f(string(k), v)

				return nil
			},
		),
	)
}
