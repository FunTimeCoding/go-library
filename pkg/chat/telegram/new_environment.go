package telegram

import (
	"github.com/funtimecoding/go-library/pkg/chat/telegram/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.TokenEnvironment),
		environment.Fallback(constant.DatabaseEnvironment, ""),
	)
}
