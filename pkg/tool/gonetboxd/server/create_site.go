package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateSite(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body server.CreateNameRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	i := s.client.MustCreateSite(body.Name)
	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(
		w,
		server.Site{Identifier: i.Identifier, Name: i.Name},
	)
}
