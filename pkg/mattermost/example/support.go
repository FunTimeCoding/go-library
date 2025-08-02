package example

import (
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/mattermost"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func Support() {
	user1 := internal.Mattermost(
		mattermost.WithToken(environment.Get("MATTERMOST_TOKEN_USER1")),
	)
	user2 := internal.Mattermost(
		mattermost.WithToken(environment.Get("MATTERMOST_TOKEN_USER2")),
	)
	h := user1.DefaultChannel()
	t := user1.PostSimple(
		h,
		"I have a problem with my server, its not responding anymore",
	)
	user2.Reply(h, t, "Can you try to restart it?")
	// TODO: wipe channel before
	// TODO: multiple different problems with successful answers
	// TODO: import all messages into Chroma
}
