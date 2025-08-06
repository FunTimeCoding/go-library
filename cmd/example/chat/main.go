package main

import (
	discord "github.com/funtimecoding/go-library/pkg/chat/discord/example"
	mattermost "github.com/funtimecoding/go-library/pkg/chat/mattermost/example"
	telegram "github.com/funtimecoding/go-library/pkg/chat/telegram/example"
)

func main() {
	if false {
		discord.DeleteLoop()
	}

	if false {
		mattermost.Dialog()
		mattermost.Support()
		mattermost.Post()
	}

	if false {
		telegram.OllamaSession()
		telegram.OllamaReply()
		telegram.Update()
		telegram.Echo()
		telegram.User()
	}
}
