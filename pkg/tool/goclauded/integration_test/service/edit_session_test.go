//go:build local

package service

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/service_tester"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service/argument/edit_session"
	"testing"
)

func TestEditSession(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("session-1")
	s.EditSession(
		"session-1",
		edit_session.New().
			WithAlias("my-project").
			WithDescription("Refactored the CLI"),
	)
	e := s.Store.GetSession("session-1")
	assert.String(t, "my-project", e.AliasValue())
	assert.String(t, "Refactored the CLI", e.Description)
}
