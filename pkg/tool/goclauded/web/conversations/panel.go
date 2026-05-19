package conversations

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) panel(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier := r.PathValue("id")
	session := s.claude.Resolve(identifier)

	if session.Identifier == "" {
		w.Header().Set(constant.ContentType, "text/html; charset=utf-8")
		errors.PanicOnError(
			html.P(gomponents.Text("Session not found.")).Render(w),
		)

		return
	}

	messages := s.claude.Messages(session.Identifier)
	var nodes []gomponents.Node

	for i := range messages {
		if messages[i].IsMeta {
			continue
		}

		nodes = append(nodes, messageBlock(&messages[i]))
	}

	w.Header().Set(constant.ContentType, "text/html; charset=utf-8")
	errors.PanicOnError(gomponents.Group(nodes).Render(w))
}
