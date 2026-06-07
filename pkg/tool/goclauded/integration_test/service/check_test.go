//go:build local

package service

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
	"time"
)

func TestCheckNewSession(t *testing.T) {
	s := service_tester.New(t)
	r := s.Check("session-1")
	assert.True(t, r.Callsign != "")
	assert.True(t, r.Changed)
}

func TestCheckNoChanges(t *testing.T) {
	s := service_tester.New(t)
	first := s.Check("session-1")
	assert.True(t, first.Changed)
	second := s.Check("session-1")
	assert.False(t, second.Changed)
}

func TestCheckWithMessage(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	s.Store.EnsureSession("session-2")
	s.Store.Advance(time.Second)
	s.Send(r1.Callsign, "", "broadcast")
	r := s.Check("session-1")
	assert.True(t, r.Changed)
	assert.Count(t, 1, r.Messages)
}

func TestCheckWithTimeout(t *testing.T) {
	s := service_tester.New(t)
	r := s.Check("session-1")
	s.Announce("session-1", r.Callsign, "working", "")
	s.Store.Advance(2 * time.Hour)
	s.Service.RunTimeoutSweep()
	r = s.Check("session-1")
	assert.True(t, r.Changed)
	assert.StringContains(t, "inactivity", r.TimeoutMessage)
}

func TestCheckWithReannounce(t *testing.T) {
	s := service_tester.New(t)
	r := s.Check("session-1")
	s.Store.BindModelContextSession(r.Callsign, "mcp-abc")
	s.Store.ClearBindings()
	r = s.Check("session-1")
	assert.True(t, r.Changed)
	assert.True(t, r.Reannounce)
}

func TestCheckSessionsPopulated(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	s.Store.Advance(time.Second)
	s.Store.Announce(r2.Callsign, "other work", "pkg/auth")
	r := s.Check("session-1")
	assert.True(t, r.Changed)
	assert.True(t, len(r.Sessions) >= 2)
	found := false

	for _, e := range r.Sessions {
		if e.CallsignValue() == r2.Callsign {
			assert.String(t, "other work", e.Topic)
			found = true
		}
	}

	assert.True(t, found)
}

func TestCheckCompleteTimeout(t *testing.T) {
	s := service_tester.New(t)
	r := s.Check("session-1")
	s.Announce("session-1", r.Callsign, "topic", "")
	s.Store.CompleteTask(r.Callsign)
	s.Store.Advance(time.Hour)
	s.Service.RunTimeoutSweep()
	r = s.Check("session-1")
	assert.True(t, r.Changed)
	assert.StringContains(
		t,
		"completing without re-announcing",
		r.TimeoutMessage,
	)
}

func TestCheckConsumesMessages(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	s.Send(r2.Callsign, r1.Callsign, "hello")
	assert.Count(
		t,
		1,
		s.Check("session-1").Messages,
	)
	assert.Count(
		t,
		0,
		s.Check("session-1").Messages,
	)
}

func TestCheckWithCompletionActivity(t *testing.T) {
	s := service_tester.New(t)
	s.Check("session-1")
	r2 := s.Store.EnsureSession("session-2")
	s.Store.Announce(r2.Callsign, "some work", "")
	s.Complete("session-2", r2.Callsign, "some work", "done")
	s.Store.Advance(time.Second)
	r := s.Check("session-1")
	assert.True(t, r.Changed)
	assert.True(t, len(r.Completions) > 0)
}
