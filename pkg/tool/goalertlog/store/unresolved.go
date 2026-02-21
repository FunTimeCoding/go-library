package store

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/bolt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"go.etcd.io/bbolt"
	"sort"
)

func (s *Store) Unresolved() []UnresolvedRecord {
	var result []UnresolvedRecord

	s.client.View(
		func(t *bbolt.Tx) error {
			b := s.client.Bucket(t, Bucket)

			if b == nil {
				return nil
			}

			bolt.For(
				b,
				func(
					k string,
					v []byte,
				) {
					var r Record
					errors.PanicOnError(json.Unmarshal(v, &r))

					if r.End == nil {
						result = append(
							result,
							UnresolvedRecord{
								Fingerprint: r.Fingerprint,
								Key:         k,
							},
						)
					}
				},
			)

			return nil
		},
	)

	sort.Slice(
		result, func(i, j int) bool {
			return result[i].Key < result[j].Key
		},
	)

	return result
}
