//go:build local

package service

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/integration_test/service_tester"
	"testing"
)

func TestPushDocumentAppearsInList(t *testing.T) {
	s := service_tester.New(t)
	assert.FatalOnError(
		t,
		s.Service.PushDocument(
			"notes",
			"first.md",
			"# First Note\n\nPushed directly via API.\n",
			"",
		),
	)
	entries, e := s.Service.ListDocuments("notes")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, entries)
	assert.String(t, "First Note", entries[0].Title)
	assert.String(t, "qmd://notes/first.md", entries[0].VirtualPath)
}

func TestPushDocumentAppearsInSearch(t *testing.T) {
	s := service_tester.New(t)
	assert.FatalOnError(
		t,
		s.Service.PushDocument(
			"notes",
			"searchable.md",
			"# Searchable\n\nThis document contains a unique keyword: platypus.\n",
			"",
		),
	)
	outcome := s.Service.Search(
		"platypus",
		10,
		"",
		false,
		"",
		"keyword",
	)
	assert.Count(t, 1, outcome.Results)
	assert.String(t, "Searchable", outcome.Results[0].Title)
}

func TestPushDocumentDedup(t *testing.T) {
	s := service_tester.New(t)
	body := "# Identical\n\nSame content pushed twice.\n"
	assert.FatalOnError(
		t,
		s.Service.PushDocument("notes", "first.md", body, ""),
	)
	assert.FatalOnError(
		t,
		s.Service.PushDocument("notes", "first.md", body, ""),
	)
	entries, e := s.Service.ListDocuments("notes")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, entries)
	status := s.Service.MustStatus()
	assert.Integer(t, 1, status.TotalDocuments)
}

func TestPushDocumentUpdate(t *testing.T) {
	s := service_tester.New(t)
	assert.FatalOnError(
		t,
		s.Service.PushDocument(
			"notes",
			"evolving.md",
			"# Original\n\nFirst version.\n",
			"",
		),
	)
	assert.FatalOnError(
		t,
		s.Service.PushDocument(
			"notes",
			"evolving.md",
			"# Updated\n\nSecond version with revised content.\n",
			"",
		),
	)
	d, _, e := s.Service.GetDocument("notes/evolving.md")
	assert.FatalOnError(t, e)
	assert.NotNil(t, d)
	assert.String(t, "Updated", d.Title)
	assert.StringContains(t, "revised content", d.Body)
}

func TestPushDocumentCreatesCollection(t *testing.T) {
	s := service_tester.New(t)
	assert.FatalOnError(
		t,
		s.Service.PushDocument(
			"dynamic",
			"auto.md",
			"# Auto\n\nCollection created on first push.\n",
			"",
		),
	)
	status := s.Service.MustStatus()
	found := false

	for _, c := range status.Collections {
		if c.Name == "dynamic" {
			found = true
			assert.Integer(t, 1, c.DocumentCount)
		}
	}

	assert.True(t, found)
}
