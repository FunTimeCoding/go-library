package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"net/http"
)

func (s *Server) PostSend(
	w http.ResponseWriter,
	r *http.Request,
) {
	var body server.SendRequest
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))
	var to string

	if body.To != nil {
		to = *body.To
	}

	s.service.Store.SendMessage(body.Name, to, body.Body)
	w.WriteHeader(http.StatusOK)
}
