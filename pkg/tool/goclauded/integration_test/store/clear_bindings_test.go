//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/store_tester"
	"testing"
)

func TestClearBindings(t *testing.T) {
	s := store_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Store.BindModelContextSession(r.Callsign, "mcp-session-abc")
	e := s.Store.GetSession("session-1")
	assert.String(t, "mcp-session-abc", e.ModelContextSession)
	s.Store.ClearBindings()
	e = s.Store.GetSession("session-1")
	assert.String(t, "", e.ModelContextSession)
	assert.True(t, e.NeedsReannounce)
}
