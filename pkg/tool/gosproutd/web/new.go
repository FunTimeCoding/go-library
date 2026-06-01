package web

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/store"
	"github.com/funtimecoding/go-library/pkg/web/layout"
	theme "github.com/funtimecoding/go-library/pkg/web/theme/constant"
	"github.com/funtimecoding/go-library/pkg/web/view"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func New(s *store.Store, n face.EventNotifier) *Server {
	return &Server{
		store:    s,
		notifier: n,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(theme.Sprout).
				WithStyle(inlineCSS).
				WithScript("https://cdn.jsdelivr.net/npm/sortablejs@1.15.6/Sortable.min.js").
				WithLiveEndpoint("/event").
				WithFooter(
					html.Script(gomponents.Raw(sortableJS)),
				),
		),
	}
}
