package store

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"go.etcd.io/bbolt"
)

func (s *Store) Save(r Record) string {
	key := fmt.Sprintf("%s|%d", r.Fingerprint, r.Start.UnixNano())
	s.client.Update(
		func(t *bbolt.Tx) error {
			v, e := json.Marshal(r)
			errors.PanicOnError(e)
			s.client.PutBytes(s.client.Bucket(t, Bucket), key, v)

			return nil
		},
	)

	return key
}
