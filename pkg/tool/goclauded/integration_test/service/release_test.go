//go:build local

package service

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
	"time"
)

func TestRelease(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Store.Announce(r.Callsign, "working", "")
	sessions := s.Store.ListSessions()
	assert.Count(t, 1, sessions)
	s.Release("session-1", r.Callsign)
	sessions = s.Store.ListSessions()
	assert.Count(t, 0, sessions)
	events := s.Store.EventsSince(
		time.Time{},
		time.Time{},
		constant.Release,
		10,
		0,
	)
	assert.Count(t, 1, events)
}
