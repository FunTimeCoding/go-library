package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
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
			Label:    constant.DashboardTitle,
			Path:     constant.DashboardPath,
			Category: "navigate",
		},
		palette.Command{
			Label:    constant.EntriesTitle,
			Path:     constant.EntriesPath,
			Category: "navigate",
		},
		palette.Command{
			Label:    constant.AddEntryTitle,
			Path:     constant.AddEntryPath,
			Category: "action",
		},
	)

	return &Server{
		store:    s,
		registry: registry,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(theme.Archive).
				WithStyle(inlineCSS).
				WithCommandPalette("/palette").
				WithItems(
					navigation_item.New(
						constant.DashboardPath,
						constant.DashboardTitle,
					),
					navigation_item.New(
						constant.EntriesPath,
						constant.EntriesTitle,
					),
					navigation_item.New(
						constant.AddEntryPath,
						constant.AddEntryTitle,
					),
				),
		),
	}
}
