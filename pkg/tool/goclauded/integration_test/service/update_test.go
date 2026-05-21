//go:build local

package service

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
	"time"
)

func TestUpdate(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Store.Announce(r.Callsign, "initial topic", "")
	s.Service.Update("session-1", r.Callsign, "new scope", "pkg/tool")
	e := s.Store.GetSession("session-1")
	assert.String(t, "new scope", e.Topic)
	assert.String(t, "pkg/tool", e.Files)
	events := s.Store.EventsSince(
		time.Time{},
		time.Time{},
		constant.Update,
		10,
		0,
	)
	assert.Count(t, 1, events)
	completions := s.Store.CompletionsBySession("session-1")
	assert.Count(t, 1, completions)
	assert.String(t, "new scope", completions[0].Topic)
}
