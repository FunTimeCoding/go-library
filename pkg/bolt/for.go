package bolt

import "go.etcd.io/bbolt"

func For(
	b *bbolt.Bucket,
	f func(
		k string,
		v []byte,
	),
) error {
	return b.ForEach(
		func(
			k []byte,
			v []byte,
		) error {
			f(string(k), v)

			return nil
		},
	)
}
