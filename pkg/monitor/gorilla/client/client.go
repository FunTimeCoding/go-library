package client

import (
	"github.com/gorilla/websocket"
	"net/url"
)

type Client struct {
	locator    url.URL
	connection *websocket.Conn
	done       chan struct{}
	receive    []string
}
