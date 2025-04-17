package telegram

import (
	"github.com/funtimecoding/go-library/pkg/bolt"
	"github.com/funtimecoding/go-library/pkg/telegram/constant"
	"github.com/funtimecoding/go-library/pkg/telegram/database/channel"
	"github.com/funtimecoding/go-library/pkg/telegram/database/user"
	"go.etcd.io/bbolt"
)

func (c *Client) readDatabase() {
	if c.database == nil {
		return
	}

	c.database.View(
		func(t *bbolt.Tx) error {
			bolt.For(
				c.database.Bucket(t, constant.ChannelBucket),
				func(
					k string,
					v []byte,
				) {
					a := channel.New()
					a.Decode(v)
					c.channels = append(c.channels, a)
				},
			)
			bolt.For(
				c.database.Bucket(t, constant.UserBucket),
				func(
					k string,
					v []byte,
				) {
					s := user.New()
					s.Decode(v)
					c.users = append(c.users, s)
				},
			)

			return nil
		},
	)
}
