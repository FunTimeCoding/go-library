package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
	"time"
)

func (s *Server) editSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	e := s.entryFromQuery(r)
	errors.PanicOnError(r.ParseForm())
	e.Action = r.FormValue("action")
	e.User = r.FormValue("user")
	e.System = r.FormValue("system")
	e.Service = r.FormValue("service")
	e.Description = r.FormValue("description")

	if v := r.FormValue("timestamp"); v != "" {
		if t, f := time.Parse("2006-01-02T15:04", v); f == nil {
			e.Timestamp = t
		}
	}

	errors.PanicOnError(s.store.Update(e))
	renderFragment(w, detailRow(e))
}
