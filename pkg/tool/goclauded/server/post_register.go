package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/sweep"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) PostRegister(
	w http.ResponseWriter,
	r *http.Request,
) {
	var body server.RegisterRequest
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))
	result := s.service.Store.EnsureSession(body.Session)
	s.service.Store.MarkNeedsRoster(body.Session)
	sweep.Run(s.harborPath)
	s.service.Store.EnrichSession(body.Session, s.claude)
	s.service.Store.LogEvent(body.Session, "register", result.Name, "", "")
	s.logger.Structured(
		"register",
		"claude_session_identifier",
		body.Session,
		constant.SessionName,
		result.Name,
		"new",
		result.New,
	)
	web.EncodeNotation(
		w,
		server.RegisterResponse{
			Name: result.Name,
		},
	)
}
