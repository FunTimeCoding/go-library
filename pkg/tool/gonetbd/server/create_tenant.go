package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateTenant(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body server.CreateNameRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	t := s.client.MustCreateTenant(body.Name)
	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(
		w,
		server.Tenant{
			Identifier: t.Identifier, Name: t.Name,
		},
	)
}
