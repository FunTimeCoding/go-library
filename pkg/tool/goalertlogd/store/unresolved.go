package store

import (
	"github.com/funtimecoding/go-library/pkg/bolt"
	"github.com/funtimecoding/go-library/pkg/notation"
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
					notation.DecodeBytesStrict(v, &r, false)

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
		result,
		func(i, j int) bool {
			return result[i].Key < result[j].Key
		},
	)

	return result
}
