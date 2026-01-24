package example

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/common"
)

func Support() {
	user1 := common.Mattermost(
		mattermost.WithToken(environment.Required("MATTERMOST_TOKEN_USER1")),
	)
	user2 := common.Mattermost(
		mattermost.WithToken(environment.Required("MATTERMOST_TOKEN_USER2")),
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
