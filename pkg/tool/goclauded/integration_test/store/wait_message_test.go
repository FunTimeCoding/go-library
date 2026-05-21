//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/store_tester"
	"testing"
	"time"
)

func TestWaitMessageNotListening(t *testing.T) {
	s := store_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	messages := s.Store.WaitMessage(r.Callsign, time.Millisecond)
	assert.True(t, messages == nil)
}

func TestWaitMessageWithPending(t *testing.T) {
	s := store_tester.New(t)
	r1 := s.Store.EnsureSession("session-1")
	r2 := s.Store.EnsureSession("session-2")
	s.Store.SetListening(r1.Callsign, true)
	s.Store.SendMessage(r2.Callsign, r1.Callsign, "wake up")
	messages := s.Store.WaitMessage(r1.Callsign, time.Second)
	assert.Count(t, 1, messages)
	assert.String(t, "wake up", messages[0].Body)
}

func TestWaitMessageTimeout(t *testing.T) {
	s := store_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Store.SetListening(r.Callsign, true)
	messages := s.Store.WaitMessage(r.Callsign, 10*time.Millisecond)
	assert.True(t, messages == nil)
}
