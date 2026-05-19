package model_context

import "github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"

type profileResponse struct {
	Always      []store.Memory        `json:"always"`
	Relevant    []store.SearchResult  `json:"relevant,omitempty"`
	Index       []store.MemorySummary `json:"index"`
	Impressions []store.Impression    `json:"impressions,omitempty"`
}
