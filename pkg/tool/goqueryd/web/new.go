package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/go-library/pkg/web/layout"
	"github.com/funtimecoding/go-library/pkg/web/layout/navigation_item"
	theme "github.com/funtimecoding/go-library/pkg/web/theme/constant"
	"github.com/funtimecoding/go-library/pkg/web/view"
)

func New(s *service.Service) *Server {
	return &Server{
		service: s,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(theme.Slate).
				WithStyle(inlineCSS).
				WithItems(
					navigation_item.New(
						constant.DashboardPath,
						constant.DashboardTitle,
					),
					navigation_item.New(
						constant.SearchPath,
						constant.SearchTitle,
					),
					navigation_item.New(
						constant.CollectionsPath,
						constant.CollectionsTitle,
					),
				),
		),
	}
}
