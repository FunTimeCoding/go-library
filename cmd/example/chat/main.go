package main

import (
	discord "github.com/funtimecoding/go-library/pkg/chat/discord/example"
	mattermost "github.com/funtimecoding/go-library/pkg/chat/mattermost/example"
	telegram "github.com/funtimecoding/go-library/pkg/chat/telegram/example"
)

func main() {
	mattermost.Before()

	if false {
		discord.DeleteLoop()

		mattermost.Dialog()
		mattermost.Support()
		mattermost.Post()

		telegram.OllamaSession()
		telegram.OllamaReply()
		telegram.Update()
		telegram.Echo()
		telegram.User()
	}
}
