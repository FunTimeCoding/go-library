package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/worker"
	"github.com/funtimecoding/go-library/pkg/web/layout"
	"github.com/funtimecoding/go-library/pkg/web/layout/navigation_item"
	theme "github.com/funtimecoding/go-library/pkg/web/theme/constant"
	"github.com/funtimecoding/go-library/pkg/web/view"
)

func New(
	s *store.Store,
	p *worker.Worker,
) *Server {
	return &Server{
		store:  s,
		worker: p,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(theme.Sentinel).
				WithStyle(inlineCSS).
				WithItems(
					navigation_item.New(
						constant.DashboardPath,
						constant.DashboardTitle,
					),
					navigation_item.New(
						constant.RecentPath,
						constant.RecentTitle,
					),
				),
		),
	}
}
