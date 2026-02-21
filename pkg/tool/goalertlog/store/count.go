package store

import (
	"github.com/funtimecoding/go-library/pkg/bolt"
	"go.etcd.io/bbolt"
)

func (s *Store) Count() int {
	var result int

	s.client.View(
		func(t *bbolt.Tx) error {
			b := s.client.Bucket(t, Bucket)

			if b == nil {
				return nil
			}

			bolt.For(
				b,
				func(
					_ string,
					_ []byte,
				) {
					result++
				},
			)

			return nil
		},
	)

	return result
}
