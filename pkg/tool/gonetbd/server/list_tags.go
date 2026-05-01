package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListTags(
	w http.ResponseWriter,
	_ *http.Request,
) {
	tags := s.client.MustTags()
	result := make([]server.Tag, 0, len(tags))

	for _, t := range tags {
		result = append(
			result,
			server.Tag{Identifier: t.Identifier, Name: t.Name},
		)
	}

	web.EncodeNotation(w, result)
}
