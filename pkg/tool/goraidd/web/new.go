package web

import (
	"github.com/funtimecoding/go-library/pkg/raid_parser"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/store"
	"github.com/funtimecoding/go-library/pkg/web/layout"
	"github.com/funtimecoding/go-library/pkg/web/layout/navigation_item"
	theme "github.com/funtimecoding/go-library/pkg/web/theme/constant"
	"github.com/funtimecoding/go-library/pkg/web/view"
)

func New(
	s *store.Store,
	elitePath string,
	outputPath string,
	p *raid_parser.Client,
) *Server {
	return &Server{
		store:      s,
		elitePath:  elitePath,
		outputPath: outputPath,
		parser:     p,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(theme.Tyria).
				WithStyle(inlineCSS).
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
