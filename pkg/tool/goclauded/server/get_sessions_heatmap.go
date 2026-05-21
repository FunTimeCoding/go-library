package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
	"sort"
)

func (s *Server) GetSessionsHeatmap(
	w http.ResponseWriter,
	_ *http.Request,
	p server.GetSessionsHeatmapParams,
) {
	sessions := s.service.Sessions()
	bash := p.Bash != nil && *p.Bash
	type stats struct {
		Calls    int
		Sessions int
	}
	counts := map[string]*stats{}
	totalCalls := 0
	sessionsWithCalls := 0

	for _, session := range sessions {
		calls := s.service.ToolCalls(session.Identifier)

		if len(calls) == 0 {
			continue
		}

		sessionsWithCalls++
		seen := map[string]bool{}

		for _, call := range calls {
			var key string

			if bash && call.Name == "Bash" {
				key = bashGroupKey(call.Detail)

				if key == "" {
					continue
				}
			} else if bash {
				continue
			} else {
				key = call.Name
			}

			totalCalls++

			if counts[key] == nil {
				counts[key] = &stats{}
			}

			counts[key].Calls++

			if !seen[key] {
				counts[key].Sessions++
				seen[key] = true
			}
		}
	}

	var entries []server.HeatmapEntry

	for name, st := range counts {
		entries = append(
			entries,
			server.HeatmapEntry{
				Name:     name,
				Calls:    st.Calls,
				Sessions: st.Sessions,
			},
		)
	}

	sort.Slice(
		entries,
		func(i, j int) bool {
			return entries[i].Calls > entries[j].Calls
		},
	)
	web.EncodeNotation(
		w,
		server.HeatmapResponse{
			SessionCount:      len(sessions),
			SessionsWithCalls: sessionsWithCalls,
			TotalCalls:        totalCalls,
			Entries:           entries,
		},
	)
}
