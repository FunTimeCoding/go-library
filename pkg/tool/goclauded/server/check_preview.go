package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
	"strings"
)

func (s *Server) checkPreview(
	w http.ResponseWriter,
	sessionIdentifier string,
) {
	name := s.service.CallsignBySessionIdentifier(sessionIdentifier)

	if name == "" {
		web.EncodeNotation(
			w,
			server.CheckResponse{
				Changed:     false,
				Sessions:    []server.SessionEntry{},
				Messages:    []server.Message{},
				Completions: []server.CompletionEntry{},
			},
		)

		return
	}

	var entries []server.SessionEntry

	for _, e := range s.service.ListSessions() {
		entry := server.SessionEntry{
			Callsign: e.CallsignValue(),
			Topic:    e.Topic,
		}

		if e.Files != "" {
			files := strings.Split(e.Files, "\n")
			entry.Files = &files
		}

		entries = append(entries, entry)
	}

	var completions []server.CompletionEntry

	for _, c := range s.service.RecentCompletions() {
		completions = append(
			completions,
			server.CompletionEntry{
				Name:  c.Name,
				Topic: c.Topic,
				Kind:  c.Kind,
			},
		)
	}

	if entries == nil {
		entries = []server.SessionEntry{}
	}

	if completions == nil {
		completions = []server.CompletionEntry{}
	}

	web.EncodeNotation(
		w,
		server.CheckResponse{
			Callsign:    name,
			Changed:     true,
			Sessions:    entries,
			Messages:    []server.Message{},
			Completions: completions,
		},
	)
}
