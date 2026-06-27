package server

import (
	"github.com/funtimecoding/go-library/pkg/face"
	proxFace "github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/face"
)

func New(
	v proxFace.Service,
	r face.Reporter,
) *Server {
	return &Server{service: v, reporter: r}
}
