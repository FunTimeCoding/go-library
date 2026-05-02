package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) add(
	w http.ResponseWriter,
	_ *http.Request,
) {
	renderPage(
		w,
		layout(
			"Add Entry",
			"/add",
			html.H1(gomponents.Text("Add Entry")),
			html.Div(html.ID("result")),
			addForm(),
		),
	)
}
