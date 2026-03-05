package store

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"go.etcd.io/bbolt"
	"time"
)

func (s *Store) Resolve(key string) {
	s.client.Update(
		func(t *bbolt.Tx) error {
			b := s.client.Bucket(t, Bucket)
			v := b.Get([]byte(key))

			if v == nil {
				return nil
			}

			var r Record
			errors.PanicOnError(json.Unmarshal(v, &r))
			r.End = new(time.Now())
			updated, e := json.Marshal(r)
			errors.PanicOnError(e)
			s.client.PutBytes(b, key, updated)

			return nil
		},
	)
}
