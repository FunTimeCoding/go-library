package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetSessionMessages(
	w http.ResponseWriter,
	r *http.Request,
	identifier string,
) {
	messages := s.service.Messages(identifier)
	var result []server.SessionMessage

	for _, m := range messages {
		entry := server.SessionMessage{
			Role:      m.Role,
			Text:      m.Text,
			Timestamp: m.Timestamp,
		}

		if m.IsMeta {
			entry.IsMeta = &m.IsMeta
		}

		result = append(result, entry)
	}

	web.EncodeNotation(
		w,
		server.MessagesResponse{Messages: result},
	)
}
