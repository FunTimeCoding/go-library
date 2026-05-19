package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"net/http"
	"strings"
)

func (s *Server) PostAnnounce(
	w http.ResponseWriter,
	r *http.Request,
) {
	var body server.AnnounceRequest
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))
	var files string

	if body.Files != nil {
		files = strings.Join(*body.Files, "\n")
	}

	s.service.Store.Announce(body.Name, body.Topic, files)
	w.WriteHeader(http.StatusOK)
}
