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
	name := s.service.Store.NameBySessionIdentifier(sessionIdentifier)

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

	for _, e := range s.service.Store.ListSessions() {
		entry := server.SessionEntry{
			Name:  e.Name,
			Topic: e.Topic,
		}

		if e.Files != "" {
			files := strings.Split(e.Files, "\n")
			entry.Files = &files
		}

		entries = append(entries, entry)
	}

	var completions []server.CompletionEntry

	for _, c := range s.service.Store.RecentCompletions() {
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
			Name:        name,
			Changed:     true,
			Sessions:    entries,
			Messages:    []server.Message{},
			Completions: completions,
		},
	)
}
