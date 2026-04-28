package store

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"go.etcd.io/bbolt"
)

func (s *Store) Save(r Record) string {
	key := fmt.Sprintf("%s|%d", r.Fingerprint, r.Start.UnixNano())
	s.client.Update(
		func(t *bbolt.Tx) error {
			s.client.PutBytes(
				s.client.Bucket(t, Bucket),
				key,
				notation.Marshal(r),
			)

			return nil
		},
	)

	return key
}
