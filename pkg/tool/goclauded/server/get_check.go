package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
	"strings"
)

func (s *Server) GetCheck(
	w http.ResponseWriter,
	_ *http.Request,
	params server.GetCheckParams,
) {
	preview := params.Preview != nil && *params.Preview

	if preview {
		s.checkPreview(w, params.Session)

		return
	}

	result := s.service.Store.EnsureSession(params.Session)
	changed := s.service.Store.HasChanges(result.Name, result.LastSeen)
	s.logger.Structured(
		"hook_check",
		"claude_session_identifier",
		params.Session,
		constant.SessionName,
		result.Name,
		"new",
		result.New,
		"changed",
		changed,
	)
	var entries []server.SessionEntry
	var messages []server.Message

	if changed {
		sessions := s.service.Store.ListSessions()

		for _, e := range sessions {
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

		pending := s.service.Store.PendingMessages(result.Name)

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
	}

	var completions []server.CompletionEntry

	if changed {
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
	}

	if messages == nil {
		messages = []server.Message{}
	}

	if entries == nil {
		entries = []server.SessionEntry{}
	}

	if completions == nil {
		completions = []server.CompletionEntry{}
	}

	var memoryActivity *[]server.MemoryActivityEntry

	if !result.LastSeen.IsZero() {
		since := result.LastSeen.UTC().Format("2006-01-02T15:04:05Z")
		activity := s.service.MemoryActivity(since)

		if len(activity) > 0 {
			var entries []server.MemoryActivityEntry

			for _, a := range activity {
				entries = append(
					entries,
					server.MemoryActivityEntry{
						MemoryIdentifier: a.MemoryIdentifier,
						Name:             a.Name,
						ChangeType:       a.ChangeType,
						Source:           a.Source,
					},
				)
			}

			memoryActivity = &entries
		}
	}

	response := server.CheckResponse{
		Name:           result.Name,
		Changed:        changed,
		Sessions:       entries,
		Messages:       messages,
		Completions:    completions,
		MemoryActivity: memoryActivity,
	}

	if memoryActivity != nil {
		response.Changed = true
	}

	if timeout := s.service.Store.ConsumeTimeout(result.Name); timeout != "" {
		response.TimeoutMessage = &timeout
		response.Changed = true
	}

	if s.service.Store.ConsumeReannounce(result.Name) {
		response.Reannounce = new(true)
		response.Changed = true
	}

	web.EncodeNotation(w, response)
}
