package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateTag(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body server.CreateNameRequest

	if e := json.NewDecoder(q.Body).Decode(&body); e != nil {
		web.InvalidRequestBody(w)

		return
	}

	t, e := s.client.CreateTag(body.Name)

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(
		w,
		server.Tag{
			Identifier: t.Identifier,
			Name:       t.Name,
		},
	)
}
