package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetSessionDetail(
	w http.ResponseWriter,
	_ *http.Request,
	identifier string,
) {
	sessionIdentifier := s.service.Store.ResolveSessionIdentifier(identifier)

	if sessionIdentifier == "" {
		sessionIdentifier = identifier
	}

	cs := s.claude.Resolve(sessionIdentifier)

	if cs.Identifier == "" {
		http.NotFound(w, nil)

		return
	}

	result := server.SessionDetailResponse{Identifier: cs.Identifier}

	if cs.Slug != "" {
		result.Slug = &cs.Slug
	}

	if cs.Timestamp != "" {
		result.Created = &cs.Timestamp
	}

	if sess := s.service.Store.GetSession(cs.Identifier); sess != nil {
		result.TurnCount = &sess.TurnCount
	}

	a := s.service.Store.GetAliasRecord(cs.Identifier)

	if a != nil {
		if a.Name != "" {
			result.Alias = &a.Name
		}

		if a.Description != "" {
			result.Description = &a.Description
		}
	}

	completions := s.service.Store.CompletionsBySession(cs.Identifier)

	if len(completions) > 0 {
		var entries []server.SessionDetailCompletion

		for _, c := range completions {
			entry := server.SessionDetailCompletion{
				Kind:  c.Kind,
				Topic: c.Topic,
			}

			if c.Summary != "" {
				entry.Summary = &c.Summary
			}

			ts := c.CreatedAt.Format("2006-01-02T15:04:05Z")
			entry.Timestamp = &ts
			entries = append(entries, entry)
		}

		result.Completions = &entries
	}

	if body := s.service.Store.SummaryBySession(cs.Identifier); body != "" {
		result.Summary = &body
	}

	web.EncodeNotation(w, result)
}
