package server

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"strings"
)

func (s *Server) checkPreview(
	sessionIdentifier string,
) (server.GetCheckResponseObject, error) {
	name := s.service.CallsignBySessionIdentifier(sessionIdentifier)

	if name == "" {
		return server.GetCheck200JSONResponse{
			Changed:     false,
			Sessions:    []server.SessionEntry{},
			Messages:    []server.Message{},
			Completions: []server.CompletionEntry{},
		}, nil
	}

	var entries []server.SessionEntry
	sessions, listError := s.service.ListSessions()

	if listError != nil {
		return server.GetCheck500JSONResponse(
			*s.captureFail(listError, constant.UnexpectedError),
		), nil
	}

	for _, e := range sessions {
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
	recent, completionError := s.service.RecentCompletions()

	if completionError != nil {
		return server.GetCheck500JSONResponse(
			*s.captureFail(completionError, constant.UnexpectedError),
		), nil
	}

	for _, c := range recent {
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

	return server.GetCheck200JSONResponse{
		Callsign:    name,
		Changed:     true,
		Sessions:    entries,
		Messages:    []server.Message{},
		Completions: completions,
	}, nil
}
