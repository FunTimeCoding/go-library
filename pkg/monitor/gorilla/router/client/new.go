package client

import (
	"github.com/gorilla/websocket"
	"net"
)

func New(
	name string,
	hostname string,
	a net.IP,
	c *websocket.Conn,
) *Client {
	return &Client{
		Name:       name,
		Hostname:   hostname,
		Address:    a,
		Connection: c,
	}
}
