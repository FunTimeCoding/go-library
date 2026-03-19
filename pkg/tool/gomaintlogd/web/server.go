package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	g "maragu.dev/gomponents"
	"net/http"
)

type Server struct {
	store *store.Store
}

func NewServer(s *store.Store) *Server {
	return &Server{store: s}
}

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc("GET /{$}", s.handleDashboard)
	m.HandleFunc("GET /entries", s.handleEntries)
	m.HandleFunc("GET /add", s.handleAddForm)
	m.HandleFunc("POST /add", s.handleAddSubmit)
	m.HandleFunc("GET /entry/detail", s.handleDetail)
	m.HandleFunc("GET /entry/edit", s.handleEditForm)
	m.HandleFunc("POST /entry/edit", s.handleEditSubmit)
	m.HandleFunc("POST /entry/delete", s.handleDelete)
}

func (s *Server) isHTMX(r *http.Request) bool {
	return r.Header.Get(constant.ExtendedRequest) == "true"
}

func renderPage(
	w http.ResponseWriter,
	page g.Node,
) {
	w.Header().Set(constant.ContentType, "text/html; charset=utf-8")
	errors.PanicOnError(page.Render(w))
}

func renderFragment(
	w http.ResponseWriter,
	fragment g.Node,
) {
	w.Header().Set(constant.ContentType, "text/html; charset=utf-8")
	errors.PanicOnError(fragment.Render(w))
}
