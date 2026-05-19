//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestHybridSearchReturnsResults(t *testing.T) {
	s, o := indexedTestStore(t)
	defer s.Close()
	_, e := s.Embed(o)
	assert.FatalOnError(t, e)
	option := NewSearchOption("search pipeline", 10)
	option.Reranker = nil
	results, e := s.SearchHybrid(option, o)
	assert.FatalOnError(t, e)
	assert.Greater(t, 0, float64(len(results)))
}

func TestHybridSearchDiffersFromKeyword(t *testing.T) {
	s, o := indexedTestStore(t)
	defer s.Close()
	_, e := s.Embed(o)
	assert.FatalOnError(t, e)
	keyword := s.MustSearchKeyword("context resolution", 5, "", false)
	option := NewSearchOption("context resolution", 5)
	option.Reranker = nil
	hybrid, e := s.SearchHybrid(option, o)
	assert.FatalOnError(t, e)

	if len(keyword) > 1 && len(hybrid) > 1 {
		keywordFirst := keyword[0].FilePath
		hybridFirst := hybrid[0].FilePath
		keywordScore := keyword[0].Score
		hybridScore := hybrid[0].Score

		if keywordFirst == hybridFirst {
			assert.True(t, keywordScore != hybridScore)
		}
	}
}

func TestSearchFallsBackToKeywordWithoutEmbeddings(t *testing.T) {
	s, o := indexedTestStore(t)
	defer s.Close()
	option := NewSearchOption("chunking strategy", 10)
	outcome := s.SearchWithFallback(option, o)
	assert.True(t, outcome.Degraded)
	assert.Greater(t, 0, float64(len(outcome.Results)))
}
