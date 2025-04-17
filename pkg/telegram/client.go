package telegram

import (
	"github.com/funtimecoding/go-library/pkg/bolt"
	"github.com/funtimecoding/go-library/pkg/telegram/database/channel"
	"github.com/funtimecoding/go-library/pkg/telegram/database/user"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Client struct {
	client  *tgbotapi.BotAPI
	Verbose bool

	database *bolt.Client
	channels []*channel.Channel
	users    []*user.User
}
