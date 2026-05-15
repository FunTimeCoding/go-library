package server

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/netbox"
)

type Server struct {
	client   *netbox.Client
	reporter face.Reporter
}
