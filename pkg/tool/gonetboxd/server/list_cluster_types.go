package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListClusterTypes(
	w http.ResponseWriter,
	_ *http.Request,
) {
	types, e := s.client.ClusterTypes()

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	result := make([]server.ClusterType, 0, len(types))

	for _, t := range types {
		result = append(
			result,
			server.ClusterType{Identifier: t.Identifier, Name: t.Name},
		)
	}

	web.EncodeNotation(w, result)
}
