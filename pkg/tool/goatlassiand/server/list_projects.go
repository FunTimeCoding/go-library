package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListProjects(
	w http.ResponseWriter,
	_ *http.Request,
) {
	web.EncodeNotation(w, convert.JiraProjects(s.jira.MustProjects()))
}
