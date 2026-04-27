package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/goatld/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (r *Router) SearchIssues(
	w http.ResponseWriter,
	_ *http.Request,
	p server.SearchIssuesParams,
) {
	var issues []*server.JiraIssue

	if p.Limit != nil {
		issues = convert.JiraIssues(
			r.jira.SearchLimit(*p.Limit, p.Query),
		)
	} else {
		issues = convert.JiraIssues(r.jira.Search(p.Query))
	}

	web.EncodeNotation(w, issues)
}
