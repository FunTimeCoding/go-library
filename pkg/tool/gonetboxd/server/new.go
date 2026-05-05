package server

import "github.com/funtimecoding/go-library/pkg/netbox"

func New(c *netbox.Client) *Server {
	return &Server{client: c}
}
