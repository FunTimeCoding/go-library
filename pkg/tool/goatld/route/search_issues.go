package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/server"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (h *Router) SearchIssues(
	w http.ResponseWriter,
	_ *http.Request,
	params server.SearchIssuesParams,
) {
	var issues []server.JiraIssue

	if params.Limit != nil {
		issues = toJiraIssues(h.jira.SearchLimit(*params.Limit, params.Query))
	} else {
		issues = toJiraIssues(h.jira.Search(params.Query))
	}

	w.Header().Set(constant.ContentType, constant.Object)
	errors.PanicOnError(json.NewEncoder(w).Encode(issues))
}
