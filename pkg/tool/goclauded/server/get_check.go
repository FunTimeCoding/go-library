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
	p server.GetCheckParams,
) {
	preview := p.Preview != nil && *p.Preview

	if preview {
		s.checkPreview(w, p.Session)

		return
	}

	r := s.service.Check(p.Session)
	s.logger.Structured(
		"hook_check",
		"claude_session_identifier",
		p.Session,
		constant.Callsign,
		r.Callsign,
		"changed",
		r.Changed,
	)
	var entries []server.SessionEntry

	for _, e := range r.Sessions {
		entry := server.SessionEntry{
			Callsign: e.CallsignValue(),
			Topic:    e.Topic,
		}

		if e.Files != "" {
			files := strings.Split(e.Files, "\n")
			entry.Files = &files
		}

		labels := s.service.LabelsBySession(e.Identifier)

		if len(labels) > 0 {
			var le []server.LabelEntry

			for _, l := range labels {
				le = append(
					le,
					server.LabelEntry{Key: l.Key, Value: l.Value},
				)
			}

			entry.Labels = &le
		}

		entries = append(entries, entry)
	}

	var messages []server.Message

	for _, m := range r.Messages {
		messages = append(
			messages,
			server.Message{
				From:      m.FromName,
				Body:      m.Body,
				Timestamp: m.CreatedAt.Format("2006-01-02T15:04:05Z"),
			},
		)
	}

	var completions []server.CompletionEntry

	for _, c := range r.Completions {
		completions = append(
			completions,
			server.CompletionEntry{
				Name:  c.Name,
				Topic: c.Topic,
				Kind:  c.Kind,
			},
		)
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

	if len(r.MemoryActivity) > 0 {
		var ma []server.MemoryActivityEntry

		for _, a := range r.MemoryActivity {
			ma = append(
				ma,
				server.MemoryActivityEntry{
					MemoryIdentifier: a.MemoryIdentifier,
					Name:             a.Name,
					ChangeType:       a.ChangeType,
					Source:           a.Source,
				},
			)
		}

		memoryActivity = &ma
	}

	response := server.CheckResponse{
		Callsign:       r.Callsign,
		Changed:        r.Changed,
		Sessions:       entries,
		Messages:       messages,
		Completions:    completions,
		MemoryActivity: memoryActivity,
	}

	if len(r.Pulses) > 0 {
		var pe []server.PulseEntry

		for _, p := range r.Pulses {
			pe = append(
				pe,
				server.PulseEntry{
					Body:      p.Body,
					Timestamp: p.CreatedAt.Format("2006-01-02T15:04:05Z"),
				},
			)
		}

		response.Pulses = &pe
	}

	if r.TimeoutMessage != "" {
		response.TimeoutMessage = &r.TimeoutMessage
	}

	if r.Reannounce {
		response.Reannounce = new(true)
	}

	web.EncodeNotation(w, response)
}
