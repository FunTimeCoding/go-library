package telegram

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/bolt"
	"github.com/funtimecoding/go-library/pkg/telegram/constant"
	"github.com/funtimecoding/go-library/pkg/telegram/database/channel"
	"github.com/opsgenie/opsgenie-go-sdk-v2/user"
	"go.etcd.io/bbolt"
)

func (c *Client) ReadDatabase() ([]*channel.Channel, []*user.User) {
	var channels []*channel.Channel
	var users []*user.User

	if c.database != nil {
		c.database.Update(
			func(t *bbolt.Tx) error {
				c.database.CreateBucket(t, constant.UserBucket)
				c.database.CreateBucket(t, constant.ChannelBucket)

				return nil
			},
		)
		c.database.Update(
			func(t *bbolt.Tx) error {
				h := c.database.Bucket(t, constant.ChannelBucket)
				test := channel.New()
				test.Name = "MyTestChannel"
				test.Identifier = 123456789
				c.database.PutBytes(h, test.Name, test.Encode())

				return nil
			},
		)
		c.database.View(
			func(t *bbolt.Tx) error {
				h := c.database.Bucket(t, constant.ChannelBucket)
				bolt.For(
					h,
					func(
						k string,
						v []byte,
					) {
						fmt.Printf("Channel: %s %v\n", k, v)
						a := channel.New()
						a.Decode(v)
						channels = append(channels, a)
					},
				)
				u := c.database.Bucket(t, constant.UserBucket)
				bolt.For(
					u,
					func(
						k string,
						v []byte,
					) {
						fmt.Printf("User: %s %v\n", k, v)
					},
				)

				return nil
			},
		)
	}

	for _, a := range channels {
		fmt.Println(a.Format())
	}

	return channels, users
}
