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
	d := s.service.SessionDetail(identifier)

	if d == nil {
		http.NotFound(w, nil)

		return
	}

	result := server.SessionDetailResponse{Identifier: d.Identifier}

	if d.Session != nil {
		if d.Session.Name != "" {
			result.Name = &d.Session.Name
		}

		if d.Session.Callsign != nil {
			result.Callsign = d.Session.Callsign
		}
	}

	if d.Slug != "" {
		result.Slug = &d.Slug
	}

	if d.Created != "" {
		result.Created = &d.Created
	}

	if d.TurnCount > 0 {
		result.TurnCount = &d.TurnCount
	}

	if d.Alias != "" {
		result.Alias = &d.Alias
	}

	if d.Description != "" {
		result.Description = &d.Description
	}

	if len(d.Completions) > 0 {
		var entries []server.SessionDetailCompletion

		for _, c := range d.Completions {
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

	if d.Summary != "" {
		result.Summary = &d.Summary
	}

	labels := s.service.LabelsBySession(d.Identifier)

	if len(labels) > 0 {
		var le []server.LabelEntry

		for _, l := range labels {
			le = append(
				le,
				server.LabelEntry{Key: l.Key, Value: l.Value},
			)
		}

		result.Labels = &le
	}

	pulses := s.service.PulsesBySession(d.Identifier)

	if len(pulses) > 0 {
		var pe []server.PulseEntry

		for _, p := range pulses {
			entry := server.PulseEntry{
				Body:      p.Body,
				Timestamp: p.CreatedAt.Format("2006-01-02T15:04:05Z"),
			}

			if p.FromName != "" {
				entry.From = &p.FromName
			}

			pe = append(pe, entry)
		}

		result.Pulses = &pe
	}

	web.EncodeNotation(w, result)
}
