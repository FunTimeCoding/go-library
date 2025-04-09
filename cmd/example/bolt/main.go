package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/bolt"
	"go.etcd.io/bbolt"
)

const (
	Bucket = "bucket"
	Key    = "key"
)

func main() {
	b := bolt.New("tmp/bolt.db")
	defer b.Close()
	b.Update(
		func(tx *bbolt.Tx) error {
			b.Put(b.CreateBucket(tx, Bucket), Key, "value")

			return nil
		},
	)
	b.View(
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
