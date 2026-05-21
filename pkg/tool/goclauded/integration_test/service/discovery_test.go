//go:build local

package service

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/session"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
)

func TestCheckConsistencyDiscoversUnregistered(t *testing.T) {
	s := service_tester.New(t)
	s.Client.AddSession(&session.Session{
		Identifier: "jsonl-only",
		Slug:       "test-session",
		Lines:      42,
	})
	s.Service.CheckConsistency()
	e := s.Store.GetSession("jsonl-only")
	assert.True(t, e != nil)
	assert.String(t, "test-session", e.Slug)
	assert.True(t, e.Lines == 42)
	assert.True(t, e.Callsign == nil)
	assert.String(t, "", e.Name)
}

func TestCheckConsistencySkipsExisting(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("registered")
	s.Client.AddSession(&session.Session{
		Identifier: "registered",
		Slug:       "already-known",
	})
	s.Service.CheckConsistency()
	e := s.Store.GetSession("registered")
	assert.True(t, e.Callsign != nil)
}

func TestBackfillSessionsEnriches(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("session-1")
	s.Client.AddSession(&session.Session{
		Identifier: "session-1",
		Slug:       "my-slug",
		Lines:      100,
		CWD:        "/home/user",
		Branch:     "main",
		Timestamp:  "2026-05-21T10:00:00Z",
	})
	s.Client.FirstMessages["session-1"] = "hello world"
	s.Client.UserMessages["session-1"] = []string{"hello world", "fix the bug"}
	s.Service.BackfillSessions()
	e := s.Store.GetSession("session-1")
	assert.String(t, "my-slug", e.Slug)
	assert.True(t, e.Lines == 100)
	assert.String(t, "/home/user", e.CWD)
	assert.String(t, "main", e.Branch)
	assert.String(t, "hello world", e.FirstMessage)
	assert.True(t, e.TurnCount == 2)
}
