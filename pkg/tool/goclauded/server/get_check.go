package server

import (
	"context"
	library "github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"strings"
)

func (s *Server) GetCheck(
	_ context.Context,
	r server.GetCheckRequestObject,
) (server.GetCheckResponseObject, error) {
	preview := r.Params.Preview != nil && *r.Params.Preview

	if preview {
		return s.checkPreview(r.Params.Session)
	}

	result, checkError := s.service.Check(r.Params.Session)

	if checkError != nil {
		return server.GetCheck500JSONResponse(
			*s.captureFail(checkError, library.UnexpectedError),
		), nil
	}

	s.logger.Structured(
		"hook_check",
		"claude_session_identifier",
		r.Params.Session,
		constant.Callsign,
		result.Callsign,
		"changed",
		result.Changed,
	)
	var entries []server.SessionEntry

	for _, e := range result.Sessions {
		entry := server.SessionEntry{
			Callsign: e.CallsignValue(),
			Topic:    e.Topic,
		}

		if e.Files != "" {
			files := strings.Split(e.Files, "\n")
			entry.Files = &files
		}

		labels, labelError := s.service.LabelsBySession(e.Identifier)

		if labelError != nil {
			return server.GetCheck500JSONResponse(
				*s.captureFail(labelError, library.UnexpectedError),
			), nil
		}

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

	for _, m := range result.Messages {
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

	for _, c := range result.Completions {
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

	if len(result.MemoryActivity) > 0 {
		var ma []server.MemoryActivityEntry

		for _, a := range result.MemoryActivity {
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
		Callsign:       result.Callsign,
		Changed:        result.Changed,
		Sessions:       entries,
		Messages:       messages,
		Completions:    completions,
		MemoryActivity: memoryActivity,
	}

	if len(result.Pulses) > 0 {
		var pe []server.PulseEntry

		for _, p := range result.Pulses {
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

	if result.TimeoutMessage != "" {
		response.TimeoutMessage = &result.TimeoutMessage
	}

	if result.Reannounce {
		response.Reannounce = new(true)
	}

	return server.GetCheck200JSONResponse(response), nil
}
