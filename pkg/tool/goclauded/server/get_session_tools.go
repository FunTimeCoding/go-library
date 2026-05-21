package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
	"sort"
)

func (s *Server) GetSessionTools(
	w http.ResponseWriter,
	r *http.Request,
	identifier string,
) {
	calls := s.service.ToolCalls(identifier)
	var entries []server.ToolCallEntry

	for _, c := range calls {
		entries = append(
			entries,
			server.ToolCallEntry{
				Name:       c.Name,
				Identifier: &c.Identifier,
				Timestamp:  c.Timestamp,
				Detail:     &c.Detail,
			},
		)
	}

	counts := map[string]int{}

	for _, c := range calls {
		counts[c.Name]++
	}

	var sorted []server.ToolCount

	for name, count := range counts {
		sorted = append(
			sorted,
			server.ToolCount{
				Name:  name,
				Count: count,
			},
		)
	}

	sort.Slice(
		sorted,
		func(i, j int) bool {
			return sorted[i].Count > sorted[j].Count
		},
	)
	web.EncodeNotation(
		w,
		server.ToolsResponse{
			Calls:  entries,
			Counts: sorted,
			Total:  len(calls),
		},
	)
}
