package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/go-library/pkg/web/palette"
	"github.com/funtimecoding/go-library/pkg/web/view"
)

type Server struct {
	service  *service.Service
	view     *view.View
	registry *palette.Registry
}
