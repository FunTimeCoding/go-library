package bolt

import "go.etcd.io/bbolt"

type Client struct {
	client *bbolt.DB
}
