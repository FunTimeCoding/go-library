package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/go-library/pkg/web/view"
)

type Server struct {
	service *service.Service
	view    *view.View
}
