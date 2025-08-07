package message

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"testing"
)

func TestMessage(t *testing.T) {
	assert.True(
		t,
		New(
			&tgbotapi.Message{From: &tgbotapi.User{UserName: strings.Alfa}},
		) != nil,
	)
}
