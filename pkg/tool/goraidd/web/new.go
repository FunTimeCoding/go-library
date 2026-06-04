package web

import (
	"github.com/funtimecoding/go-library/pkg/raid_parser"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/store"
	"github.com/funtimecoding/go-library/pkg/web/layout"
	"github.com/funtimecoding/go-library/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/go-library/pkg/web/palette"
	theme "github.com/funtimecoding/go-library/pkg/web/theme/constant"
	"github.com/funtimecoding/go-library/pkg/web/view"
)

func New(
	s *store.Store,
	elitePath string,
	outputPath string,
	p *raid_parser.Client,
) *Server {
	registry := palette.NewRegistry()
	registry.Register(
		palette.Command{
			Label:    constant.LogsTitle,
			Path:     constant.LogsPath,
			Category: "navigate",
		},
		palette.Command{
			Label:    constant.ReportsTitle,
			Path:     constant.ReportsPath,
			Category: "navigate",
		},
		palette.Command{
			Label:    constant.RaidsTitle,
			Path:     constant.RaidsPath,
			Category: "navigate",
		},
		palette.Command{
			Label:    constant.PlayersTitle,
			Path:     constant.PlayersPath,
			Category: "navigate",
		},
	)

	return &Server{
		store:      s,
		elitePath:  elitePath,
		outputPath: outputPath,
		parser:     p,
		registry:   registry,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(theme.Tyria).
				WithStyle(inlineCSS).
				WithCommandPalette("/palette").
				WithItems(
					navigation_item.New(
						constant.LogsPath,
						constant.LogsTitle,
					),
					navigation_item.New(
						constant.ReportsPath,
						constant.ReportsTitle,
					),
					navigation_item.New(
						constant.RaidsPath,
						constant.RaidsTitle,
					),
					navigation_item.New(
						constant.PlayersPath,
						constant.PlayersTitle,
					),
				),
		),
	}
}
