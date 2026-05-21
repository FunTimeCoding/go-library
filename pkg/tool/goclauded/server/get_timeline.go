package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service/argument"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetTimeline(
	w http.ResponseWriter,
	_ *http.Request,
	p server.GetTimelineParams,
) {
	a := argument.Timeline{Limit: 50}

	if p.Limit != nil {
		a.Limit = *p.Limit
	}

	if p.Offset != nil {
		a.Offset = *p.Offset
	}

	if p.Since != nil {
		a.Since = *p.Since
	}

	merged := s.service.Timeline(a)
	var result []server.TimelineEntry

	for _, e := range merged {
		entry := server.TimelineEntry{
			Timestamp: e.Timestamp,
			Kind:      e.Kind,
			Actor:     e.Actor,
			Subject:   e.Subject,
		}

		if e.Detail != "" {
			entry.Detail = &e.Detail
		}

		result = append(result, entry)
	}

	if result == nil {
		result = []server.TimelineEntry{}
	}

	web.EncodeNotation(w, result)
}
