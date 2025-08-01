package telegram

import (
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/telegram/constant"
)

func NewEnvironment() *Client {
	return New(
		environment.Get(constant.TokenEnvironment),
		environment.GetDefault(constant.DatabaseEnvironment, ""),
	)
}
