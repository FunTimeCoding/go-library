package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
	"time"
)

func (s *Server) GetWait(
	w http.ResponseWriter,
	_ *http.Request,
	p server.GetWaitParams,
) {
	timeout := 30

	if p.Timeout != nil {
		timeout = *p.Timeout
	}

	pending := s.service.WaitMessage(
		p.Callsign,
		time.Duration(timeout)*time.Second,
	)
	var messages []server.Message

	for _, m := range pending {
		messages = append(
			messages,
			server.Message{
				From:      m.FromName,
				Body:      m.Body,
				Timestamp: m.CreatedAt.Format("2006-01-02T15:04:05Z"),
			},
		)
	}

	if messages == nil {
		messages = []server.Message{}
	}

	web.EncodeNotation(
		w,
		server.WaitResponse{
			Messages: messages,
		},
	)
}
