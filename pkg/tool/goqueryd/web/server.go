package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/web/search_cache"
	"github.com/funtimecoding/go-library/pkg/web/palette"
	"github.com/funtimecoding/go-library/pkg/web/view"
)

type Server struct {
	service  *service.Service
	view     *view.View
	registry *palette.Registry
	cache    *search_cache.Cache
}
