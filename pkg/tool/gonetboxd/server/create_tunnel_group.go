package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateTunnelGroup(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body server.CreateNameRequest

	if e := json.NewDecoder(q.Body).Decode(&body); e != nil {
		web.InvalidRequestBody(w)

		return
	}

	g, e := s.client.CreateTunnelGroup(body.Name)

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(
		w,
		server.TunnelGroup{
			Identifier: g.Identifier,
			Name:       g.Name,
		},
	)
}
