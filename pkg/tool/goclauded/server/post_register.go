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
	sweep.Run(s.harborPath)
	result := s.service.Register(body.Session)
	s.logger.Structured(
		"register",
		"claude_session_identifier",
		body.Session,
		constant.SessionName,
		result.Callsign,
		"new",
		result.New,
	)
	web.EncodeNotation(
		w,
		server.RegisterResponse{
			Callsign: result.Callsign,
		},
	)
}
