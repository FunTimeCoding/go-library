//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestDeleteDocument(t *testing.T) {
	s, o := openTestStore(t)
	defer s.Close()
	assert.FatalOnError(
		t,
		s.PushDocument(
			"notes",
			"doomed.md",
			"# Doomed\n\nThis document will be deleted.\n",
			"",
			o,
		),
	)
	deleted, e := s.DeleteDocument("notes", "doomed.md")
	assert.FatalOnError(t, e)
	assert.True(t, deleted)
	entries := s.MustListDocuments("notes")
	assert.Count(t, 0, entries)
}

func TestDeleteDocumentNotFound(t *testing.T) {
	s, _ := openTestStore(t)
	defer s.Close()
	deleted, e := s.DeleteDocument("notes", "nonexistent.md")
	assert.FatalOnError(t, e)
	assert.False(t, deleted)
}

func TestDeleteDocumentCleansOrphanedContent(t *testing.T) {
	s, o := openTestStore(t)
	defer s.Close()
	assert.FatalOnError(
		t,
		s.PushDocument(
			"notes",
			"unique.md",
			"# Unique\n\nOnly one document uses this content.\n",
			"",
			o,
		),
	)
	_, e := s.DeleteDocument("notes", "unique.md")
	assert.FatalOnError(t, e)
	assert.Integer(t, 0, s.ContentCount())
}

func TestDeleteDocumentPreservesSharedContent(t *testing.T) {
	s, o := openTestStore(t)
	defer s.Close()
	body := "# Shared\n\nTwo documents share this body.\n"
	assert.FatalOnError(t, s.PushDocument("first", "shared.md", body, "", o))
	assert.FatalOnError(t, s.PushDocument("second", "shared.md", body, "", o))
	_, e := s.DeleteDocument("first", "shared.md")
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, s.ContentCount())
	d := s.MustGetDocument("second/shared.md")
	assert.NotNil(t, d)
}

func TestDeleteDocumentRemovesFromSearch(t *testing.T) {
	s, o := openTestStore(t)
	defer s.Close()
	assert.FatalOnError(
		t,
		s.PushDocument(
			"notes",
			"findable.md",
			"# Findable\n\nContains the keyword wolverine.\n",
			"",
			o,
		),
	)
	results := s.MustSearchKeyword("wolverine", 10, "", false)
	assert.Count(t, 1, results)
	_, e := s.DeleteDocument("notes", "findable.md")
	assert.FatalOnError(t, e)
	results = s.MustSearchKeyword("wolverine", 10, "", false)
	assert.Count(t, 0, results)
}
