package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/server"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (h *Router) GetTransitions(
	w http.ResponseWriter,
	_ *http.Request,
	key string,
) {
	transitions := h.jira.Transitions(key)
	result := make([]server.JiraTransition, 0, len(transitions))

	for _, t := range transitions {
		entry := server.JiraTransition{
			Identifier: t.ID,
			Name:       t.Name,
		}

		if t.To.Name != "" {
			entry.ToStatus = &t.To.Name
		}

		result = append(result, entry)
	}

	w.Header().Set(constant.ContentType, constant.Object)
	errors.PanicOnError(json.NewEncoder(w).Encode(result))
}
