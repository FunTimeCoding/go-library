package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func New(token string) *Client {
	client, e := discordgo.New(fmt.Sprintf("Bot %s", token))
	errors.PanicOnError(e)

	if false {
		fmt.Printf("Identify: %+v\n", client.Identify)
		fmt.Printf("State: %+v\n", client.State)
	}

	return &Client{client: client}
}
