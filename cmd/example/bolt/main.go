package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/bolt"
	"go.etcd.io/bbolt"
)

func main() {
	b := bolt.New("tmp/bolt.db")
	defer b.Close()
	b.MustUpdate(
		func(t *bbolt.Tx) error {
			b.MustPut(b.MustCreateBucket(t, Bucket), Key, "value")

			return nil
		},
	)
	b.MustView(
		func(t *bbolt.Tx) error {
			bucket := b.Bucket(t, Bucket)

			if bucket == nil {
				return nil
			}

			fmt.Println(b.Get(bucket, Key))

			return nil
		},
	)
}
