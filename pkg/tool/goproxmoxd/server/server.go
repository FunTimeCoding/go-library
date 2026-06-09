package server

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/service"
)

type Server struct {
	service  *service.Service
	reporter face.Reporter
}
