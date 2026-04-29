package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goatld/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetIssue(
	w http.ResponseWriter,
	_ *http.Request,
	key string,
) {
	web.EncodeNotation(w, convert.JiraIssue(s.jira.Issue(key)))
}
