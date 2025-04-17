package telegram

import (
	"github.com/funtimecoding/go-library/pkg/bolt"
	"github.com/funtimecoding/go-library/pkg/telegram/constant"
	"go.etcd.io/bbolt"
)

func (c *Client) setDatabase(b *bolt.Client) *Client {
	c.database = b
	c.database.Update(
		func(t *bbolt.Tx) error {
			c.database.CreateBucket(t, constant.UserBucket)
			c.database.CreateBucket(t, constant.ChannelBucket)

			return nil
		},
	)
	c.readDatabase()

	return c
}
