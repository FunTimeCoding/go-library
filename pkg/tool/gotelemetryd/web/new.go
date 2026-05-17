package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/store"
	"github.com/funtimecoding/go-library/pkg/web/layout"
	"github.com/funtimecoding/go-library/pkg/web/layout/navigation_item"
	theme "github.com/funtimecoding/go-library/pkg/web/theme/constant"
	"github.com/funtimecoding/go-library/pkg/web/view"
)

func New(s *store.Store) *Server {
	return &Server{
		store: s,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(theme.Archive).
				WithItems(
					navigation_item.New(
						constant.HeatmapPath,
						constant.HeatmapTitle,
					),
					navigation_item.New(
						constant.EventsPath,
						constant.EventsTitle,
					),
				),
		),
	}
}
