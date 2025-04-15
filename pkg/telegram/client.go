package telegram

import (
	"github.com/funtimecoding/go-library/pkg/bolt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Client struct {
	client   *tgbotapi.BotAPI
	Verbose  bool
	database *bolt.Client
}
