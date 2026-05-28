//go:build local

package service

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/integration_test/service_tester"
	"testing"
)

func TestSearchKeyword(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	outcome := s.Service.Search(
		"hybrid search pipeline",
		10,
		"",
		false,
		"",
		"keyword",
	)
	assert.Greater(t, 0, float64(len(outcome.Results)))
	assert.String(t, "Search Pipeline", outcome.Results[0].Title)
}

func TestSearchHybrid(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	_, e := s.Service.Embed()
	assert.FatalOnError(t, e)
	outcome := s.Service.Search(
		"semantic meaning",
		10,
		"",
		false,
		"",
		"hybrid",
	)
	assert.Greater(t, 0, float64(len(outcome.Results)))
}

func TestSearchCollectionFilter(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	outcome := s.Service.Search(
		"chunking",
		10,
		"test",
		false,
		"",
		"keyword",
	)
	assert.Greater(t, 0, float64(len(outcome.Results)))
	assert.String(t, "test", outcome.Results[0].Collection)
}

func TestSearchDegradedWithoutEmbeddings(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	outcome := s.Service.Search(
		"chunking strategy",
		10,
		"",
		false,
		"",
		"hybrid",
	)
	assert.True(t, outcome.Degraded)
	assert.Greater(t, 0, float64(len(outcome.Results)))
}
