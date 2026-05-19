package service

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service/argument"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/timeline"
	"strings"
	"time"
)

func (s *Service) Timeline(q argument.Timeline) []*timeline.Entry {
	fetch := q.Limit + q.Offset
	var since, before time.Time

	if q.Since != "" {
		t, e := time.Parse(time.RFC3339, q.Since)

		if e == nil {
			since = t
		}
	}

	if q.Before != "" {
		t, e := time.Parse(time.RFC3339, q.Before)

		if e == nil {
			before = t
		}
	}

	var entries []*timeline.Entry

	for _, e := range s.Store.EventsSince(
		since,
		before,
		q.Kind,
		fetch,
		0,
	) {
		entries = append(
			entries,
			timeline.FromEvent(
				e.Kind,
				e.Name,
				e.Body,
				e.Target,
				e.CreatedAt,
				q.Full,
			),
		)
	}

	if q.Kind == "" || strings.HasPrefix(q.Kind, "memory_") {
		entries = append(entries, s.FetchVersions(q.Since, fetch)...)
	}

	return timeline.Merge(entries, q.Limit, q.Offset)
}
