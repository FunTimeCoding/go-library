package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goatld/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/generated/server"
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
			s.jira.SearchLimit(*p.Limit, p.Query),
		)
	} else {
		issues = convert.JiraIssues(s.jira.Search(p.Query))
	}

	web.EncodeNotation(w, issues)
}
