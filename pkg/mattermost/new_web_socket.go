package mattermost

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost-server/v6/model"
)

func newWebSocket(
	host string,
	token string,
) *model.WebSocketClient {
	result, e := model.NewWebSocketClient4(
		fmt.Sprintf("wss://%s", host),
		token,
	)
	errors.PanicOnError(e)

	return result
}
