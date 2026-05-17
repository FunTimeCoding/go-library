package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
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
				WithStyle(inlineCSS).
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
