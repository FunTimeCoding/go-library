package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/goatld/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) GetTransitions(
	w http.ResponseWriter,
	_ *http.Request,
	key string,
) {
	transitions := h.jira.Transitions(key)
	result := make([]*server.JiraTransition, 0, len(transitions))

	for _, t := range transitions {
		r := &server.JiraTransition{Identifier: t.ID, Name: t.Name}

		if t.To.Name != "" {
			r.ToStatus = &t.To.Name
		}

		result = append(result, r)
	}

	web.EncodeNotation(w, result)
}
