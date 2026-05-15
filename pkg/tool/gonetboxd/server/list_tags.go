package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListTags(
	w http.ResponseWriter,
	_ *http.Request,
) {
	tags, e := s.client.Tags()

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	result := make([]server.Tag, 0, len(tags))

	for _, t := range tags {
		result = append(
			result,
			server.Tag{Identifier: t.Identifier, Name: t.Name},
		)
	}

	web.EncodeNotation(w, result)
}
