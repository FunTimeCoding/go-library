package telegram

import (
	"github.com/funtimecoding/go-library/pkg/bolt"
	"github.com/funtimecoding/go-library/pkg/chat/telegram/constant"
	"go.etcd.io/bbolt"
)

func (c *Client) setDatabase(b *bolt.Client) *Client {
	c.database = b
	c.database.MustUpdate(
		func(t *bbolt.Tx) error {
			c.database.MustCreateBucket(t, constant.UserBucket)
			c.database.MustCreateBucket(t, constant.ChannelBucket)

			return nil
		},
	)
	c.readDatabase()

	return c
}
