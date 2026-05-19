//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestEmbedCreatesEmbeddings(t *testing.T) {
	s, o := indexedTestStore(t)
	defer s.Close()
	result, e := s.Embed(o)
	assert.FatalOnError(t, e)
	assert.Greater(t, 0, float64(result.Documents))
	assert.Greater(t, 0, float64(result.Chunks))
	status := s.MustStatus()
	assert.Greater(t, 0, float64(status.TotalEmbeddings))
	assert.Integer(t, 0, status.PendingEmbeddings)
}

func TestEmbedIdempotent(t *testing.T) {
	s, o := indexedTestStore(t)
	defer s.Close()
	first, e := s.Embed(o)
	assert.FatalOnError(t, e)
	second, e := s.Embed(o)
	assert.FatalOnError(t, e)
	assert.Integer(t, 0, second.Documents)
	assert.Integer(t, 0, second.Chunks)
	status := s.MustStatus()
	assert.Integer(t, first.Chunks, status.TotalEmbeddings)
}

func TestEmbedAfterPush(t *testing.T) {
	s, o := openTestStore(t)
	defer s.Close()
	assert.FatalOnError(
		t,
		s.PushDocument(
			"notes",
			"embed-me.md",
			"# Embed Test\n\nThis document should get embedded.\n",
			"",
			o,
		),
	)
	status := s.MustStatus()
	assert.Greater(t, 0, float64(status.TotalEmbeddings))
}
