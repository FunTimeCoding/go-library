package server

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/netbox"
)

func New(
	c *netbox.Client,
	r face.Reporter,
) *Server {
	return &Server{client: c, reporter: r}
}
