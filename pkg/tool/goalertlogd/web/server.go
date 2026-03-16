package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/poller"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	g "maragu.dev/gomponents"
	"net/http"
)

type Server struct {
	store  *store.Store
	poller *poller.Poller
}

func NewServer(
	s *store.Store,
	p *poller.Poller,
) *Server {
	return &Server{
		store:  s,
		poller: p,
	}
}

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc("GET /{$}", s.handleDashboard)
	m.HandleFunc("GET /recent", s.handleRecent)
	m.HandleFunc("GET /alerts", s.handleAlerts)
}

func (s *Server) isHTMX(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}

func renderPage(w http.ResponseWriter, page g.Node) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	errors.PanicOnError(page.Render(w))
}

func renderFragment(w http.ResponseWriter, fragment g.Node) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	errors.PanicOnError(fragment.Render(w))
}
