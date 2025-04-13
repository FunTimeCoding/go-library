package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/mattermost/mattermost-server/v6/model"
)

func main() {
	host := environment.Get("MATTERMOST_HOST", 1)
	token := environment.Get("MATTERMOST_TOKEN", 1)
	teamName := environment.Get("MATTERMOST_TEAM", 1)
	channelName := environment.Get("MATTERMOST_CHANNEL", 1)
	client := model.NewAPIv4Client(fmt.Sprintf("https://%s", host))
	client.SetOAuthToken(token)
	t, _, teamFail := client.GetTeamByName(teamName, "")
	errors.PanicOnError(teamFail)
	c, _, channelFail := client.GetChannelByName(channelName, t.Id, "")
	errors.PanicOnError(channelFail)
	_, _, postFail := client.CreatePost(
		&model.Post{ChannelId: c.Id, Message: "Hello friend"},
	)
	errors.PanicOnError(postFail)
}
