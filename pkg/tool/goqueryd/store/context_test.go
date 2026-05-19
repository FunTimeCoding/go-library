//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestContextHierarchicalResolution(t *testing.T) {
	s, _ := indexedTestStore(t)
	defer s.Close()
	s.AddContext("test", "/", "root context")
	s.AddContext("test", "/tools/", "tools context")
	root := s.resolveContext("test", "alpha.md")
	assert.String(t, "root context", root)
	sub := s.resolveContext("test", "tools/gamma.md")
	assert.StringContains(t, "root context", sub)
	assert.StringContains(t, "tools context", sub)
}

func TestContextAttachedToSearchResults(t *testing.T) {
	s, _ := indexedTestStore(t)
	defer s.Close()
	s.AddContext("test", "/", "all documents")
	results := s.MustSearchKeyword("hybrid search pipeline", 10, "", false)
	assert.Count(t, 1, results)
	assert.String(t, "all documents", results[0].Context)
}

func TestContextAddOverwrites(t *testing.T) {
	s, _ := openTestStore(t)
	defer s.Close()
	directory := t.TempDir()
	s.AddCollection("test", directory, "**/*.md")
	s.AddContext("test", "/", "first")
	s.AddContext("test", "/", "second")
	entries := s.ListContexts()
	assert.Count(t, 1, entries)
	assert.String(t, "second", entries[0].Description)
}

func TestContextRemove(t *testing.T) {
	s, _ := openTestStore(t)
	defer s.Close()
	directory := t.TempDir()
	s.AddCollection("test", directory, "**/*.md")
	s.AddContext("test", "/", "to remove")
	removed := s.RemoveContext("test", "/")
	assert.True(t, removed)
	entries := s.ListContexts()
	assert.Count(t, 0, entries)
}

func TestContextRemoveNotFound(t *testing.T) {
	s, _ := openTestStore(t)
	defer s.Close()
	removed := s.RemoveContext("nonexistent", "/")
	assert.False(t, removed)
}
