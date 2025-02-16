package client

import (
	"github.com/gorilla/websocket"
	"net"
)

type Client struct {
	Name       string
	Hostname   string
	Address    net.IP
	Connection *websocket.Conn
}
