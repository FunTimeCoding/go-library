package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service/argument"
	"net/http"
)

func (s *Server) PostEditSession(
	w http.ResponseWriter,
	r *http.Request,
) {
	var b server.EditSessionRequest
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&b))
	a := argument.NewEditSession()
	a.Alias = b.Name
	a.Description = b.Description
	a.Topic = b.Topic
	a.Files = b.Files
	s.service.EditSession(b.Session, a)
	w.WriteHeader(http.StatusOK)
}
