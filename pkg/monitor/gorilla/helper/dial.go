package helper

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/gorilla/websocket"
	"net/url"
)

func Dial(u url.URL) *websocket.Conn {
	result, _, e := websocket.DefaultDialer.Dial(u.String(), nil)
	errors.PanicOnError(e)

	return result
}
