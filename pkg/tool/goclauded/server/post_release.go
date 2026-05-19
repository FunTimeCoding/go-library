package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"net/http"
)

func (s *Server) PostRelease(
	w http.ResponseWriter,
	r *http.Request,
) {
	var body server.ReleaseRequest
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))
	s.service.Store.ReleaseByName(body.Name)
	w.WriteHeader(http.StatusOK)
}
