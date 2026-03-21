package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) addSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(r.ParseForm())
	e := &store.Entry{
		Action:      r.FormValue("action"),
		User:        r.FormValue("user"),
		System:      r.FormValue("system"),
		Service:     r.FormValue("service"),
		Description: r.FormValue("description"),
	}
	errors.PanicOnError(s.store.Add(e))
	renderFragment(
		w,
		h.Div(
			h.Div(
				h.Class("success"),
				g.Textf(
					"Entry added: %s by %s",
					e.Action,
					e.User,
				),
			),
			addForm(),
		),
	)
}
