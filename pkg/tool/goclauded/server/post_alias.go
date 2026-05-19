package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"net/http"
)

func (s *Server) PostAlias(
	w http.ResponseWriter,
	r *http.Request,
) {
	var body server.AliasRequest
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))

	if body.Name != nil {
		s.service.Store.SetAlias(body.Session, *body.Name)
	}

	if body.Description != nil {
		s.service.Store.SetDescription(body.Session, *body.Description)
	}

	w.WriteHeader(http.StatusOK)
}
