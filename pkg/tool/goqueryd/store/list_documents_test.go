//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestListDocuments(t *testing.T) {
	s, _ := indexedTestStore(t)
	defer s.Close()
	entries := s.MustListDocuments("test")
	assert.Count(t, 5, entries)
	assert.String(t, "qmd://test/alpha.md", entries[0].VirtualPath)
}

func TestListDocumentsEmptyCollection(t *testing.T) {
	s, _ := openTestStore(t)
	defer s.Close()
	entries := s.MustListDocuments("nonexistent")
	assert.Count(t, 0, entries)
}
