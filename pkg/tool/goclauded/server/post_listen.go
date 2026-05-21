package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"net/http"
)

func (s *Server) PostListen(
	w http.ResponseWriter,
	r *http.Request,
) {
	var body server.ListenRequest
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))
	listening := body.Listening != nil && *body.Listening
	s.service.SetListening(body.Callsign, listening)
	w.WriteHeader(http.StatusOK)
}
