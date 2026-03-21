package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	g "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	h "maragu.dev/gomponents/html"
	"net/http"
	"strconv"
	"time"
)

func (s *Server) handleDetail(
	w http.ResponseWriter,
	r *http.Request,
) {
	e := s.entryFromQuery(r)
	renderFragment(w, detailRow(e))
}

func (s *Server) handleEditForm(
	w http.ResponseWriter,
	r *http.Request,
) {
	e := s.entryFromQuery(r)
	renderFragment(w, editForm(e))
}

func (s *Server) handleEditSubmit(
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

func (s *Server) handleDelete(
	w http.ResponseWriter,
	r *http.Request,
) {
	id, e := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)
	errors.PanicOnError(e)
	errors.PanicOnError(s.store.Delete(uint(id)))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func (s *Server) entryFromQuery(r *http.Request) *store.Entry {
	id, e := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)
	errors.PanicOnError(e)
	entry, f := s.store.Get(uint(id))
	errors.PanicOnError(f)

	return entry
}

func detailRow(e *store.Entry) g.Node {
	target := fmt.Sprintf("#detail-%d", e.ID)

	return h.Tr(
		h.ID(fmt.Sprintf("detail-%d", e.ID)),
		h.Class("detail-row"),
		h.Td(
			g.Attr("colspan", "6"),
			h.Div(
				h.Class("detail-content"),
				h.Div(
					h.Class("grid"),
					h.Div(
						h.Strong(g.Text("Timestamp: ")),
						g.Text(formatTime(e.Timestamp)),
					),
					h.Div(
						h.Strong(g.Text("Action: ")),
						g.Text(e.Action),
					),
					h.Div(
						h.Strong(g.Text("User: ")),
						g.Text(e.User),
					),
				),
				h.Div(
					h.Class("grid"),
					h.Div(
						h.Strong(g.Text("System: ")),
						g.Text(e.System),
					),
					h.Div(
						h.Strong(g.Text("Service: ")),
						g.Text(e.Service),
					),
					h.Div(),
				),
				h.Div(
					h.Strong(g.Text("Description:")),
					h.Pre(g.Text(e.Description)),
				),
				h.Div(
					h.Class("detail-actions"),
					h.Button(
						h.Class("outline"),
						hx.Get(fmt.Sprintf("/entry/edit?id=%d", e.ID)),
						hx.Target(target),
						hx.Swap("outerHTML"),
						g.Text("Edit"),
					),
					h.Button(
						h.Class("outline contrast"),
						hx.Post(fmt.Sprintf("/entry/delete?id=%d", e.ID)),
						hx.Confirm("Delete this entry?"),
						g.Attr("hx-on::after-request",
							fmt.Sprintf(
								"document.getElementById('row-%d')?.remove();document.getElementById('detail-%d')?.remove()",
								e.ID, e.ID,
							),
						),
						g.Text("Delete"),
					),
					h.Button(
						h.Type("button"),
						h.Class("outline secondary"),
						g.Attr("onclick",
							fmt.Sprintf(
								"var r=document.getElementById('detail-%d');r.style.display='none';r.innerHTML='';r.className=''",
								e.ID,
							),
						),
						g.Text("Close"),
					),
				),
			),
		),
	)
}

func editForm(e *store.Entry) g.Node {
	target := fmt.Sprintf("#detail-%d", e.ID)

	return h.Tr(
		h.ID(fmt.Sprintf("detail-%d", e.ID)),
		h.Class("detail-row"),
		h.Td(
			g.Attr("colspan", "6"),
			h.Form(
				h.Class("edit-form"),
				hx.Post(fmt.Sprintf("/entry/edit?id=%d", e.ID)),
				hx.Target(target),
				hx.Swap("outerHTML"),
				h.Div(
					h.Class("grid"),
					h.Label(
						g.Text("Action"),
						h.Input(
							h.Type("text"),
							h.Name("action"),
							h.Value(e.Action),
							h.Required(),
						),
					),
					h.Label(
						g.Text("User"),
						h.Input(
							h.Type("text"),
							h.Name("user"),
							h.Value(e.User),
							h.Required(),
						),
					),
					h.Label(
						g.Text("Timestamp"),
						h.Input(
							h.Type("datetime-local"),
							h.Name("timestamp"),
							h.Value(
								e.Timestamp.Format("2006-01-02T15:04"),
							),
						),
					),
				),
				h.Div(
					h.Class("grid"),
					h.Label(
						g.Text("System"),
						h.Input(
							h.Type("text"),
							h.Name("system"),
							h.Value(e.System),
						),
					),
					h.Label(
						g.Text("Service"),
						h.Input(
							h.Type("text"),
							h.Name("service"),
							h.Value(e.Service),
						),
					),
				),
				h.Label(
					g.Text("Description"),
					h.Textarea(
						h.Name("description"),
						h.Rows("4"),
						g.Text(e.Description),
					),
				),
				h.Div(
					h.Class("detail-actions"),
					h.Button(h.Type("submit"), g.Text("Save")),
					h.Button(
						h.Type("button"),
						h.Class("outline secondary"),
						hx.Get(fmt.Sprintf("/entry/detail?id=%d", e.ID)),
						hx.Target(target),
						hx.Swap("outerHTML"),
						g.Text("Cancel"),
					),
				),
			),
		),
	)
}
