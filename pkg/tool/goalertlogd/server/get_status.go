package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetStatus(
	w http.ResponseWriter,
	_ *http.Request,
) {
	result := server.StatusResponse{
		TotalRecords: s.store.MustCount(),
	}
	lastPoll := s.poller.LastPoll()

	if !lastPoll.IsZero() {
		result.LastPoll = new(lastPoll.Format(DateFormat))
	}

	web.EncodeNotation(w, result)
}
