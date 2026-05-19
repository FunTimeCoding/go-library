package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) profile(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	topic := q.GetString(constant.Topic, "")
	always, e := s.listMemoriesWithContent(constant.AlwaysTag)

	if e != nil {
		return s.captureFail(e, "failed to load always-tagged memories")
	}

	var relevant []store.SearchResult

	if topic != "" {
		relevant, e = s.service.Store.SearchMemories(topic, 10, "", "")

		if e != nil {
			return s.captureFail(e, "failed to search for relevant memories")
		}
	}

	alwaysIDs := map[int64]bool{}

	for _, m := range always {
		alwaysIDs[m.Identifier] = true
	}

	relevantIDs := map[int64]bool{}

	for _, r := range relevant {
		relevantIDs[r.Identifier] = true
	}

	allMemories, e := s.service.Store.ListMemories("", "", true)

	if e != nil {
		return s.captureFail(e, "failed to list memories")
	}

	var index []store.MemorySummary

	for _, m := range allMemories {
		if !alwaysIDs[m.Identifier] && !relevantIDs[m.Identifier] {
			index = append(index, m)
		}
	}

	since := time.Now().Add(-48 * time.Hour).UTC().Format(time.RFC3339)
	impressions, e := s.service.Store.RecentImpressions(since)

	if e != nil {
		impressions = nil
	}

	return response.Success(
		notation.MarshalIndent(
			profileResponse{
				Always:      always,
				Relevant:    relevant,
				Index:       index,
				Impressions: impressions,
			},
		),
	)
}
