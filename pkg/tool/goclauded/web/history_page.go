package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service/argument"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"strconv"
)

func (s *Server) historyPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	page := 1

	if v := r.URL.Query().Get("page"); v != "" {
		if n, e := strconv.Atoi(v); e == nil && n > 0 {
			page = n
		}
	}

	limit := 10
	offset := (page - 1) * limit
	entries, e := s.service.Timeline(
		argument.Timeline{
			Limit:  limit + 1,
			Offset: offset,
		},
	)
	errors.PanicOnError(e)
	hasMore := len(entries) > limit

	if hasMore {
		entries = entries[:limit]
	}

	var rows []gomponents.Node

	for _, e := range entries {
		rows = append(rows, timelineRow(e))
	}

	var content []gomponents.Node
	content = append(content, html.H3(gomponents.Text(constant.HistoryTitle)))

	if len(rows) == 0 {
		content = append(
			content,
			html.P(gomponents.Text("No events recorded.")),
		)
	} else {
		content = append(
			content,
			html.Table(
				html.THead(
					html.Tr(
						html.Th(gomponents.Text("Time")),
						html.Th(gomponents.Text("Event")),
					),
				),
				html.TBody(rows...),
			),
		)
	}

	var navigation []gomponents.Node

	if page > 1 {
		navigation = append(
			navigation,
			html.A(
				gomponents.Attr("href", paginationLink(page-1)),
				gomponents.Text("← Newer"),
			),
		)
	}

	if hasMore {
		if len(navigation) > 0 {
			navigation = append(navigation, gomponents.Text(" · "))
		}

		navigation = append(
			navigation,
			html.A(
				gomponents.Attr("href", paginationLink(page+1)),
				gomponents.Text("Older →"),
			),
		)
	}

	if len(navigation) > 0 {
		content = append(content, html.P(gomponents.Group(navigation)))
	}

	s.view.RenderPage(
		w,
		constant.HistoryTitle,
		constant.HistoryPath,
		content...,
	)
}
