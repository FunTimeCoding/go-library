package server

import (
	"github.com/funtimecoding/go-library/pkg/face"
	proxFace "github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/face"
)

type Server struct {
	service  proxFace.Service
	reporter face.Reporter
}
