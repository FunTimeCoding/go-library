package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/goatld/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) GetIssue(
	w http.ResponseWriter,
	_ *http.Request,
	key string,
) {
	web.EncodeNotation(w, convert.JiraIssue(h.jira.Issue(key)))
}
