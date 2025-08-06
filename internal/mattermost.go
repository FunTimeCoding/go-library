package internal

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/user_map"
)

func Mattermost(o ...mattermost.OptionFunc) *mattermost.Client {
	return mattermost.NewEnvironment(
		append(
			o,
			mattermost.WithUserMap(
				user_map.New(
					func(s string) string {
						return s
					},
					[]string{},
				),
			),
		)...,
	)
}
