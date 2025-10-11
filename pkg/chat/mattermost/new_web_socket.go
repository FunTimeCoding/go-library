package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"github.com/mattermost/mattermost/server/public/model"
)

func newWebSocket(
	host string,
	token string,
) *model.WebSocketClient {
	result, e := model.NewWebSocketClient4(
		locator.New(host).Scheme(constant.SecureSocket).String(),
		token,
	)
	errors.PanicOnError(e)

	return result
}
