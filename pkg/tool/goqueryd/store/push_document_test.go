//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestPushDocumentAppearsInList(t *testing.T) {
	s, o := openTestStore(t)
	defer s.Close()
	assert.FatalOnError(
		t,
		s.PushDocument(
			"notes",
			"first.md",
			"# First Note\n\nPushed directly via API.\n",
			"",
			o,
		),
	)
	entries := s.MustListDocuments("notes")
	assert.Count(t, 1, entries)
	assert.String(t, "First Note", entries[0].Title)
	assert.String(t, "qmd://notes/first.md", entries[0].VirtualPath)
}

func TestPushDocumentAppearsInSearch(t *testing.T) {
	s, o := openTestStore(t)
	defer s.Close()
	assert.FatalOnError(
		t,
		s.PushDocument(
			"notes",
			"searchable.md",
			"# Searchable\n\nThis document contains a unique keyword: platypus.\n",
			"",
			o,
		),
	)
	results := s.MustSearchKeyword("platypus", 10, "", false)
	assert.Count(t, 1, results)
	assert.String(t, "Searchable", results[0].Title)
}

func TestPushDocumentDedup(t *testing.T) {
	s, o := openTestStore(t)
	defer s.Close()
	body := "# Identical\n\nSame content pushed twice.\n"
	assert.FatalOnError(t, s.PushDocument("notes", "first.md", body, "", o))
	assert.FatalOnError(t, s.PushDocument("notes", "first.md", body, "", o))
	entries := s.MustListDocuments("notes")
	assert.Count(t, 1, entries)
	status := s.MustStatus()
	assert.Integer(t, 1, status.TotalDocuments)
}

func TestPushDocumentUpdate(t *testing.T) {
	s, o := openTestStore(t)
	defer s.Close()
	assert.FatalOnError(
		t,
		s.PushDocument(
			"notes",
			"evolving.md",
			"# Original\n\nFirst version.\n",
			"",
			o,
		),
	)
	assert.FatalOnError(
		t,
		s.PushDocument(
			"notes",
			"evolving.md",
			"# Updated\n\nSecond version with revised content.\n",
			"",
			o,
		),
	)
	d := s.MustGetDocument("notes/evolving.md")
	assert.NotNil(t, d)
	assert.String(t, "Updated", d.Title)
	assert.StringContains(t, "revised content", d.Body)
}

func TestPushDocumentCreatesCollection(t *testing.T) {
	s, o := openTestStore(t)
	defer s.Close()
	assert.FatalOnError(
		t,
		s.PushDocument(
			"dynamic",
			"auto.md",
			"# Auto\n\nCollection created on first push.\n",
			"",
			o,
		),
	)
	status := s.MustStatus()
	found := false

	for _, c := range status.Collections {
		if c.Name == "dynamic" {
			found = true
			assert.Integer(t, 1, c.DocumentCount)
		}
	}

	assert.True(t, found)
}
