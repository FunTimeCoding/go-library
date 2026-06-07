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
	r := s.EnsureSession("session-1")
	assert.FatalOnError(t, s.Store.MarkNeedsRoster("session-1"))
	first, e := s.Store.ConsumeRoster(r.Callsign)
	assert.FatalOnError(t, e)
	assert.True(t, first)
	second, e := s.Store.ConsumeRoster(r.Callsign)
	assert.FatalOnError(t, e)
	assert.False(t, second)
}

func TestHasChangesOtherSessionUpdated(t *testing.T) {
	s := store_tester.New(t)
	r1 := s.EnsureSession("session-1")
	r2 := s.EnsureSession("session-2")
	before := time.Now()
	s.Advance(time.Second)
	s.Announce(r2.Callsign, "new topic", "")
	changed, e := s.Store.HasChanges(r1.Callsign, before)
	assert.FatalOnError(t, e)
	assert.True(t, changed)
}

func TestHasChangesUnreadMessage(t *testing.T) {
	s := store_tester.New(t)
	r1 := s.EnsureSession("session-1")
	r2 := s.EnsureSession("session-2")
	s.SendMessage(r2.Callsign, r1.Callsign, "hello")
	changed, e := s.Store.HasChanges(r1.Callsign, time.Time{})
	assert.FatalOnError(t, e)
	assert.True(t, changed)
}

func TestHasChangesNoChanges(t *testing.T) {
	s := store_tester.New(t)
	r := s.EnsureSession("session-1")
	changed, e := s.Store.HasChanges(r.Callsign, time.Now())
	assert.FatalOnError(t, e)
	assert.False(t, changed)
}
