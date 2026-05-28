package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) searchPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	query := r.URL.Query().Get("query")
	collection := r.URL.Query().Get(constant.Collection)
	mode := r.URL.Query().Get(constant.Mode)

	if mode == "" {
		mode = "hybrid"
	}

	var content []gomponents.Node
	content = append(
		content,
		html.H3(gomponents.Text(constant.SearchTitle)),
		html.Form(
			gomponents.Attr("method", "GET"),
			gomponents.Attr("action", constant.SearchPath),
			html.Input(
				html.Type(constant.Search),
				html.Name("query"),
				gomponents.Attr("placeholder", "search documents..."),
				gomponents.Attr("value", query),
				gomponents.Attr("autocomplete", "off"),
			),
			html.Button(
				html.Type("submit"),
				gomponents.Text(constant.SearchTitle),
			),
		),
	)

	if query != "" {
		outcome := s.service.Search(
			query,
			20,
			collection,
			false,
			"",
			mode,
		)

		if outcome.Degraded {
			content = append(
				content,
				html.P(
					html.Small(
						html.Em(
							gomponents.Text("degraded to keyword search"),
						),
					),
				),
			)
		}

		if len(outcome.Results) == 0 {
			content = append(
				content,
				html.P(gomponents.Text("No results.")),
			)
		} else {
			var rows []gomponents.Node

			for _, m := range outcome.Results {
				rows = append(
					rows,
					html.Tr(
						html.Td(
							html.A(
								gomponents.Attr(
									"href",
									fmt.Sprintf(
										"/documents/%s",
										documentPath(m.VirtualPath),
									),
								),
								gomponents.Text(m.Title),
							),
						),
						html.Td(
							html.Small(
								gomponents.Text(m.VirtualPath),
							),
						),
						html.Td(
							html.Small(
								gomponents.Text(m.Snippet),
							),
						),
						html.Td(
							gomponents.Textf("%.2f", m.Score),
						),
					),
				)
			}

			content = append(
				content,
				html.P(
					gomponents.Textf(
						"%d results",
						len(outcome.Results),
					),
				),
				html.Table(
					html.THead(
						html.Tr(
							html.Th(gomponents.Text("Title")),
							html.Th(gomponents.Text("Path")),
							html.Th(gomponents.Text("Snippet")),
							html.Th(gomponents.Text("Score")),
						),
					),
					html.TBody(rows...),
				),
			)
		}
	}

	s.view.RenderPage(
		w,
		constant.SearchTitle,
		constant.SearchPath,
		content...,
	)
}
