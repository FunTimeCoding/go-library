package web

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) add(
	w http.ResponseWriter,
	_ *http.Request,
) {
	renderPage(
		w,
		layout(
			"Add Entry", "/add",
			h.H1(g.Text("Add Entry")),
			h.Div(h.ID("result")),
			addForm(),
		),
	)
}
