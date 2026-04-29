package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goatld/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetTransitions(
	w http.ResponseWriter,
	_ *http.Request,
	key string,
) {
	transitions := s.jira.Transitions(key)
	result := make([]*server.JiraTransition, 0, len(transitions))

	for _, t := range transitions {
		a := &server.JiraTransition{Identifier: t.ID, Name: t.Name}

		if t.To.Name != "" {
			a.ToStatus = &t.To.Name
		}

		result = append(result, a)
	}

	web.EncodeNotation(w, result)
}
