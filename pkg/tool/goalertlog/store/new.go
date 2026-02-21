package store

import (
	"github.com/funtimecoding/go-library/pkg/bolt"
	"go.etcd.io/bbolt"
)

func New(path string) *Store {
	c := bolt.New(path)
	c.Update(
		func(t *bbolt.Tx) error {
			c.CreateBucket(t, Bucket)

			return nil
		},
	)

	return &Store{client: c}
}
