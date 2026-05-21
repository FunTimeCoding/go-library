//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/store_tester"
	"testing"
)

func TestConsumeReannounceWhenSet(t *testing.T) {
	s := store_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Store.BindModelContextSession(r.Callsign, "mcp-session-abc")
	s.Store.ClearBindings()
	assert.True(t, s.Store.ConsumeReannounce(r.Callsign))
	assert.False(t, s.Store.ConsumeReannounce(r.Callsign))
}

func TestConsumeReannounceWhenNotSet(t *testing.T) {
	s := store_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	assert.False(t, s.Store.ConsumeReannounce(r.Callsign))
}
