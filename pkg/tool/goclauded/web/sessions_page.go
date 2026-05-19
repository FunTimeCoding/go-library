package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) sessionsPage(
	w http.ResponseWriter,
	_ *http.Request,
) {
	sessions := s.service.Store.AllSessions()
	active := s.service.Store.ListSessions()
	activeNames := make(map[string]bool, len(active))

	for _, a := range active {
		activeNames[a.Name] = true
	}

	aliases := make(map[string]string, len(sessions))

	for _, e := range sessions {
		if a := s.service.Store.GetAlias(e.Identifier); a != "" {
			aliases[e.Name] = a
		}
	}

	var rows []gomponents.Node

	for i := range sessions {
		rows = append(
			rows,
			sessionRow(&sessions[i], activeNames, aliases),
		)
	}

	var content []gomponents.Node
	content = append(
		content,
		html.H3(gomponents.Text(constant.SessionsTitle)),
	)

	if len(rows) == 0 {
		content = append(content, html.P(gomponents.Text("No sessions.")))
	} else {
		content = append(
			content,
			html.Table(
				html.THead(
					html.Tr(
						html.Th(gomponents.Text("Name")),
						html.Th(gomponents.Text("Topic")),
						html.Th(gomponents.Text("Alias")),
						html.Th(gomponents.Text("Turns")),
						html.Th(gomponents.Text("Last Seen")),
						html.Th(gomponents.Text("First Message")),
					),
				),
				html.TBody(rows...),
			),
		)
	}

	s.view.RenderPage(
		w,
		constant.SessionsTitle,
		constant.SessionsPath,
		content...,
	)
}
