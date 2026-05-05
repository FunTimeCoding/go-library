package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) SearchIssues(
	w http.ResponseWriter,
	_ *http.Request,
	p server.SearchIssuesParams,
) {
	var issues []*server.JiraIssue

	if p.Limit != nil {
		issues = convert.JiraIssues(
			s.jira.MustSearchLimit(*p.Limit, p.Query),
		)
	} else {
		issues = convert.JiraIssues(s.jira.MustSearch(p.Query))
	}

	web.EncodeNotation(w, issues)
}
