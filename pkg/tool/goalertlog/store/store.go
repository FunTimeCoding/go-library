package store

import "github.com/funtimecoding/go-library/pkg/bolt"

type Store struct {
	client *bolt.Client
}
