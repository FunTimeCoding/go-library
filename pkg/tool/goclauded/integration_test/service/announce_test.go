//go:build local

package service

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
	"time"
)

func TestAnnounce(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Service.Announce("session-1", r.Callsign, "fixing auth", "pkg/auth")
	e := s.Store.GetSession("session-1")
	assert.String(t, "fixing auth", e.Topic)
	assert.String(t, "pkg/auth", e.Files)
	events := s.Store.EventsSince(
		time.Time{},
		time.Time{},
		constant.Announce,
		10,
		0,
	)
	assert.Count(t, 1, events)
	assert.String(t, "fixing auth", events[0].Body)
}

func TestAnnounceUpdatesTopic(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Service.Announce("session-1", r.Callsign, "first topic", "")
	s.Service.Announce("session-1", r.Callsign, "second topic", "")
	assert.String(
		t,
		"second topic",
		s.Store.GetSession("session-1").Topic,
	)
}
