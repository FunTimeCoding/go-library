package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateClusterType(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body server.CreateNameRequest

	if e := json.NewDecoder(q.Body).Decode(&body); e != nil {
		web.InvalidRequestBody(w)

		return
	}

	t, e := s.client.CreateClusterType(body.Name)

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(
		w,
		server.ClusterType{
			Identifier: t.Identifier,
			Name:       t.Name,
		},
	)
}
