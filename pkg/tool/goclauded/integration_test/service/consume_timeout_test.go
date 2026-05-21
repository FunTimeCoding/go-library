//go:build local

package service

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
	"time"
)

func TestConsumeTimeoutInactivity(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Store.Announce(r.Callsign, "working on auth", "")
	s.Advance(2 * time.Hour)
	s.Service.RunTimeoutSweep()
	msg := s.Store.ConsumeTimeout(r.Callsign)
	assert.StringContains(t, "inactivity", msg)
	assert.StringContains(t, "working on auth", msg)
	assert.String(t, "", s.Store.ConsumeTimeout(r.Callsign))
}

func TestConsumeTimeoutComplete(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Store.Announce(r.Callsign, "topic", "")
	s.Store.CompleteTask(r.Callsign)
	s.Advance(time.Hour)
	s.Service.RunTimeoutSweep()
	msg := s.Store.ConsumeTimeout(r.Callsign)
	assert.StringContains(t, "completing without re-announcing", msg)
}

func TestConsumeTimeoutNone(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	assert.String(t, "", s.Store.ConsumeTimeout(r.Callsign))
}
