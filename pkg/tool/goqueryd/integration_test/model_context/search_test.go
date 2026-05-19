//go:build local

package model_context

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/integration_test/model_context_tester"
	"testing"
)

func indexFixtures(t *testing.T) *model_context_tester.Tester {
	t.Helper()
	s := model_context_tester.New(t)
	s.IndexFixtures()

	return s
}

func TestSearchKeyword(t *testing.T) {
	s := indexFixtures(t)
	result := s.SearchKeyword("hybrid search pipeline")
	assert.StringContains(t, "Search Pipeline", result)
}

func TestSearchKeywordCollectionFilter(t *testing.T) {
	s := indexFixtures(t)
	result := s.MustCallTool(
		constant.Search,
		map[string]any{
			"query":             "chunking",
			constant.Mode:       "keyword",
			constant.Collection: "test",
		},
	)
	assert.StringContains(t, "Chunking Strategy", result)
}
