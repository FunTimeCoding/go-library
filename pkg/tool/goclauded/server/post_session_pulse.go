package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"net/http"
)

func (s *Server) PostSessionPulse(
	w http.ResponseWriter,
	r *http.Request,
	identifier string,
) {
	var body server.PulseRequest
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))
	s.service.SendPulse(identifier, "", body.Body)
	w.WriteHeader(http.StatusOK)
}
