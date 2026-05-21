package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"net/http"
)

func (s *Server) PostAnnounce(
	w http.ResponseWriter,
	r *http.Request,
) {
	var body server.AnnounceRequest
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))
	var files string

	if body.Files != nil {
		files = join.NewLine(*body.Files)
	}

	identifier := s.service.ResolveByCallsign(body.Callsign)
	s.service.Announce(identifier, body.Callsign, body.Topic, files)
	w.WriteHeader(http.StatusOK)
}
