//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/store_tester"
	"testing"
	"time"
)

func TestConsumeRoster(t *testing.T) {
	s := store_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Store.MarkNeedsRoster("session-1")
	assert.True(t, s.Store.ConsumeRoster(r.Callsign))
	assert.False(t, s.Store.ConsumeRoster(r.Callsign))
}

func TestHasChangesOtherSessionUpdated(t *testing.T) {
	s := store_tester.New(t)
	r1 := s.Store.EnsureSession("session-1")
	r2 := s.Store.EnsureSession("session-2")
	before := time.Now()
	s.Advance(time.Second)
	s.Store.Announce(r2.Callsign, "new topic", "")
	assert.True(t, s.Store.HasChanges(r1.Callsign, before))
}

func TestHasChangesUnreadMessage(t *testing.T) {
	s := store_tester.New(t)
	r1 := s.Store.EnsureSession("session-1")
	r2 := s.Store.EnsureSession("session-2")
	s.Store.SendMessage(r2.Callsign, r1.Callsign, "hello")
	assert.True(t, s.Store.HasChanges(r1.Callsign, time.Time{}))
}

func TestHasChangesNoChanges(t *testing.T) {
	s := store_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	assert.False(t, s.Store.HasChanges(r.Callsign, time.Now()))
}
