package server

import "github.com/funtimecoding/go-library/pkg/proxmox"

func New(c *proxmox.Client) *Server {
	return &Server{client: c}
}
