package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/store"
	"github.com/funtimecoding/go-library/pkg/web/layout"
	"github.com/funtimecoding/go-library/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/go-library/pkg/web/palette"
	theme "github.com/funtimecoding/go-library/pkg/web/theme/constant"
	"github.com/funtimecoding/go-library/pkg/web/view"
)

func New(s *store.Store) *Server {
	registry := palette.NewRegistry()
	registry.Register(
		palette.Command{
			Label:    constant.HeatmapTitle,
			Path:     constant.HeatmapPath,
			Category: "navigate",
		},
		palette.Command{
			Label:    constant.EventsTitle,
			Path:     constant.EventsPath,
			Category: "navigate",
		},
	)

	return &Server{
		store:    s,
		registry: registry,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(theme.Archive).
				WithCommandPalette("/palette").
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
