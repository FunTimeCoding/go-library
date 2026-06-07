//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/store_tester"
	"testing"
)

func TestConsumeReannounceWhenSet(t *testing.T) {
	s := store_tester.New(t)
	r := s.EnsureSession("session-1")
	s.BindModelContextSession(r.Callsign, "mcp-session-abc")
	s.ClearBindings()
	first, e := s.Store.ConsumeReannounce(r.Callsign)
	assert.FatalOnError(t, e)
	assert.True(t, first)
	second, e := s.Store.ConsumeReannounce(r.Callsign)
	assert.FatalOnError(t, e)
	assert.False(t, second)
}

func TestConsumeReannounceWhenNotSet(t *testing.T) {
	s := store_tester.New(t)
	r := s.EnsureSession("session-1")
	result, e := s.Store.ConsumeReannounce(r.Callsign)
	assert.FatalOnError(t, e)
	assert.False(t, result)
}
