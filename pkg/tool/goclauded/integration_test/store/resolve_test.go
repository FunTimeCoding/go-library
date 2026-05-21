//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/store_tester"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service/argument"
	"testing"
)

func TestResolveByName(t *testing.T) {
	s := store_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	resolved := s.Store.ResolveByCallsign(r.Callsign)
	assert.String(t, "session-1", resolved)
}

func TestResolveByNameNotFound(t *testing.T) {
	s := store_tester.New(t)
	assert.String(t, "", s.Store.ResolveByCallsign("nonexistent"))
}

func TestResolveAlias(t *testing.T) {
	s := store_tester.New(t)
	s.Store.EnsureSession("session-1")
	edit(s, "session-1", &argument.EditSession{Alias: new("my-alias")})
	resolved := s.Store.ResolveAlias("my-alias")
	assert.String(t, "session-1", resolved)
}

func TestResolveAliasNotFound(t *testing.T) {
	s := store_tester.New(t)
	assert.String(t, "", s.Store.ResolveAlias("nonexistent"))
}

func TestSessionByName(t *testing.T) {
	s := store_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	e := s.Store.SessionByCallsign(r.Callsign)
	assert.String(t, "session-1", e.Identifier)
}

func TestSessionByNameNotFound(t *testing.T) {
	assert.True(
		t,
		store_tester.New(t).Store.SessionByCallsign("nonexistent") == nil,
	)
}
