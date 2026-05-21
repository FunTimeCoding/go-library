//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/store_tester"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service/argument"
	"testing"
)

func edit(
	s *store_tester.Tester,
	id string,
	a *argument.EditSession,
) {
	s.Store.EditSession(id, a)
}

func TestSetAndGetAlias(t *testing.T) {
	s := store_tester.New(t)
	s.Store.EnsureSession("session-1")
	edit(s, "session-1", &argument.EditSession{Alias: new("my-project")})
	assert.String(t, "my-project", s.Store.GetSession("session-1").AliasValue())
}

func TestSetAliasOverwrite(t *testing.T) {
	s := store_tester.New(t)
	s.Store.EnsureSession("session-1")
	edit(s, "session-1", &argument.EditSession{Alias: new("first-name")})
	edit(s, "session-1", &argument.EditSession{Alias: new("second-name")})
	assert.String(
		t,
		"second-name",
		s.Store.GetSession("session-1").AliasValue(),
	)
}

func TestSetAndGetDescription(t *testing.T) {
	s := store_tester.New(t)
	s.Store.EnsureSession("session-1")
	edit(
		s,
		"session-1",
		&argument.EditSession{Description: new("Fixed the auth bug")},
	)
	assert.String(
		t,
		"Fixed the auth bug",
		s.Store.GetSession("session-1").Description,
	)
}

func TestSetDescriptionWithoutAlias(t *testing.T) {
	s := store_tester.New(t)
	s.Store.EnsureSession("session-1")
	edit(
		s,
		"session-1",
		&argument.EditSession{Description: new("Standalone description")},
	)
	e := s.Store.GetSession("session-1")
	assert.String(t, "", e.AliasValue())
	assert.String(t, "Standalone description", e.Description)
}

func TestSetAliasAndDescriptionIndependently(t *testing.T) {
	s := store_tester.New(t)
	s.Store.EnsureSession("session-1")
	edit(s, "session-1", &argument.EditSession{Alias: new("my-project")})
	edit(
		s,
		"session-1",
		&argument.EditSession{Description: new("Refactored the CLI")},
	)
	e := s.Store.GetSession("session-1")
	assert.String(t, "my-project", e.AliasValue())
	assert.String(t, "Refactored the CLI", e.Description)
}

func TestSetBothAtOnce(t *testing.T) {
	s := store_tester.New(t)
	s.Store.EnsureSession("session-1")
	edit(
		s, "session-1", &argument.EditSession{
			Alias:       new("my-project"),
			Description: new("Refactored the CLI"),
		},
	)
	e := s.Store.GetSession("session-1")
	assert.String(t, "my-project", e.AliasValue())
	assert.String(t, "Refactored the CLI", e.Description)
}

func TestSetDescriptionDoesNotClearAlias(t *testing.T) {
	s := store_tester.New(t)
	s.Store.EnsureSession("session-1")
	edit(s, "session-1", &argument.EditSession{Alias: new("my-project")})
	edit(
		s,
		"session-1",
		&argument.EditSession{Description: new("New description")},
	)
	assert.String(
		t,
		"my-project",
		s.Store.GetSession("session-1").AliasValue(),
	)
}

func TestSetAliasDoesNotClearDescription(t *testing.T) {
	s := store_tester.New(t)
	s.Store.EnsureSession("session-1")
	edit(
		s,
		"session-1",
		&argument.EditSession{Description: new("Important work")},
	)
	edit(s, "session-1", &argument.EditSession{Alias: new("renamed")})
	assert.String(
		t,
		"Important work",
		s.Store.GetSession("session-1").Description,
	)
}

func TestEditTopic(t *testing.T) {
	s := store_tester.New(t)
	s.Store.EnsureSession("session-1")
	edit(
		s,
		"session-1",
		&argument.EditSession{Topic: new("debugging auth")},
	)
	assert.String(
		t,
		"debugging auth",
		s.Store.GetSession("session-1").Topic,
	)
}

func TestNoOpEdit(t *testing.T) {
	s := store_tester.New(t)
	s.Store.EnsureSession("session-1")
	edit(s, "session-1", &argument.EditSession{Alias: new("before")})
	edit(s, "session-1", &argument.EditSession{})
	assert.String(
		t,
		"before",
		s.Store.GetSession("session-1").AliasValue(),
	)
}

func TestResolveByAlias(t *testing.T) {
	s := store_tester.New(t)
	s.Store.EnsureSession("session-1")
	edit(s, "session-1", &argument.EditSession{Alias: new("my-project")})
	resolved := s.Store.ResolveSessionIdentifier("my-project")
	assert.True(t, resolved.Found())
	assert.String(t, "session-1", resolved.Identifier())
}

func TestResolveByIdentifier(t *testing.T) {
	s := store_tester.New(t)
	s.Store.EnsureSession("session-1")
	resolved := s.Store.ResolveSessionIdentifier("session-1")
	assert.True(t, resolved.Found())
	assert.String(t, "session-1", resolved.Identifier())
}

func TestAliasedSessions(t *testing.T) {
	s := store_tester.New(t)
	s.Store.EnsureSession("session-1")
	s.Store.EnsureSession("session-2")
	s.Store.EnsureSession("session-3")
	edit(s, "session-1", &argument.EditSession{Alias: new("first")})
	edit(s, "session-2", &argument.EditSession{Alias: new("second")})
	aliased := s.Store.AliasedSessions()
	assert.Count(t, 2, aliased)
}
