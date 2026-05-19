package web

import (
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/notifier"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/web/conversations"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/layout"
	"github.com/funtimecoding/go-library/pkg/web/layout/navigation_item"
	theme "github.com/funtimecoding/go-library/pkg/web/theme/constant"
	"github.com/funtimecoding/go-library/pkg/web/view"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func New(
	s *service.Service,
	n *notifier.Notifier,
	c *claude.Client,
) *Server {
	return &Server{
		service:       s,
		notifier:      n,
		conversations: conversations.New(c, s.Store),
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(theme.Hearth).
				WithStyle(inlineCSS).
				WithScript(web.ServerSide).
				WithBrandNode(
					html.Strong(
						html.Span(
							html.ID("sse-dot"),
							html.Class(
								"status-dot status-disconnected",
							),
						),
						gomponents.Text(constant.Identity.Name()),
					),
				).
				WithItems(
					navigation_item.New(
						constant.DashboardPath,
						constant.DashboardTitle,
					),
					navigation_item.New(
						constant.SessionsPath,
						constant.SessionsTitle,
					),
					navigation_item.New(
						constant.MessagesPath,
						constant.MessagesTitle,
					),
					navigation_item.New(
						constant.HistoryPath,
						constant.HistoryTitle,
					),
					navigation_item.NewExternal(
						constant.ConversationsPath,
						constant.ConversationsTitle,
					),
				).
				WithFooter(
					html.Script(
						gomponents.Raw(connectionIndicatorJS),
					),
				),
		),
	}
}
