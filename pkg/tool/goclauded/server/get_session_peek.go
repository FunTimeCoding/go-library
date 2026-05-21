package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetSessionPeek(
	w http.ResponseWriter,
	r *http.Request,
	identifier string,
) {
	p := s.service.Peek(identifier)
	web.EncodeNotation(
		w,
		server.PeekResponse{
			LineCount:    p.LineCount,
			UserMessages: p.UserMessages,
		},
	)
}
