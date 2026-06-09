package server

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/service"
)

func New(
	v *service.Service,
	r face.Reporter,
) *Server {
	return &Server{service: v, reporter: r}
}
