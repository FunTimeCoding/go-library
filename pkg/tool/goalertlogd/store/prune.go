package store

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"go.etcd.io/bbolt"
	"time"
)

func (s *Store) Prune(cutoff time.Time) int {
	var result int

	s.client.Update(
		func(t *bbolt.Tx) error {
			b := s.client.Bucket(t, Bucket)

			if b == nil {
				return nil
			}

			var keys [][]byte
			c := b.Cursor()

			for k, v := c.First(); k != nil; k, v = c.Next() {
				var r Record
				errors.PanicOnError(json.Unmarshal(v, &r))

				if r.Start.Before(cutoff) {
					keys = append(keys, k)
				}
			}

			for _, k := range keys {
				errors.PanicOnError(b.Delete(k))
			}

			result = len(keys)

			return nil
		},
	)

	return result
}
