package example

import (
	"github.com/funtimecoding/go-library/pkg/mattermost"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func Support() {
	user1 := mattermost.NewEnvironment(
		mattermost.WithToken(
			environment.Get(
				"MATTERMOST_TOKEN_USER1",
				1,
			),
		),
	)
	user2 := mattermost.NewEnvironment(
		mattermost.WithToken(
			environment.Get(
				"MATTERMOST_TOKEN_USER2",
				1,
			),
		),
	)
	h := user1.DefaultChannel()
	t := user1.PostSimple(
		h,
		"I have a problem with my server, its not responding anymore",
	)
	user2.Reply(h, t, "Can you try to restart it?")
}
