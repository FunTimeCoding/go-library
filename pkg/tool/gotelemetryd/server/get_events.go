package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/store"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetEvents(
	w http.ResponseWriter,
	_ *http.Request,
	params server.GetEventsParams,
) {
	o := store.NewQueryOption()

	if params.Tool != nil {
		o.Tool = *params.Tool
	}

	if params.Surface != nil {
		o.Surface = *params.Surface
	}

	if params.Actor != nil {
		o.Actor = *params.Actor
	}

	if params.Since != nil {
		o.Since = *params.Since
	}

	if params.Until != nil {
		o.Until = *params.Until
	}

	if params.Limit != nil {
		o.Limit = *params.Limit
	}

	if params.Offset != nil {
		o.Offset = *params.Offset
	}

	events := s.store.Recent(o)
	var entries []server.EventEntry

	for _, e := range events {
		entry := server.EventEntry{
			Id:        int(e.ID),
			Tool:      e.Tool,
			Surface:   e.Surface,
			Actor:     e.Actor,
			Outcome:   e.Outcome,
			CreatedAt: e.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}

		if e.DurationMillisecond > 0 {
			ms := int(e.DurationMillisecond)
			entry.DurationMs = &ms
		}

		entries = append(entries, entry)
	}

	if entries == nil {
		entries = []server.EventEntry{}
	}

	web.EncodeNotation(w, entries)
}
