package store

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/bolt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"go.etcd.io/bbolt"
	"sort"
	"time"
)

func (s *Store) ByTimeRange(
	start time.Time,
	end time.Time,
) []Record {
	var result []Record

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
					v []byte,
				) {
					var r Record
					errors.PanicOnError(json.Unmarshal(v, &r))

					if !r.Start.Before(start) && r.Start.Before(end) {
						result = append(result, r)
					}
				},
			)

			return nil
		},
	)

	sort.Slice(
		result, func(i, j int) bool {
			return result[i].Start.Before(result[j].Start)
		},
	)

	return result
}
