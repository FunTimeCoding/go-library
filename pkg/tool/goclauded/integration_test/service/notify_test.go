//go:build local

package service

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/service_tester"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service/argument"
	"testing"
)

func TestAnnounceNotifies(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Notifier.Reset()
	s.Announce("session-1", r.Callsign, "working", "")
	assert.True(t, s.Notifier.Notified == 1)
}

func TestCompleteNotifies(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Store.Announce(r.Callsign, "topic", "")
	s.Notifier.Reset()
	s.Complete("session-1", r.Callsign, "topic", "done")
	assert.True(t, s.Notifier.Notified == 1)
}

func TestReleaseNotifies(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Notifier.Reset()
	s.Release("session-1", r.Callsign)
	assert.True(t, s.Notifier.Notified == 1)
}

func TestSendNotifies(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Notifier.Reset()
	s.Send(r.Callsign, "", "hello")
	assert.True(t, s.Notifier.Notified == 1)
}

func TestEditSessionNotifies(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("session-1")
	s.Notifier.Reset()
	alias := "my-project"
	s.EditSession("session-1", &argument.EditSession{Alias: &alias})
	assert.True(t, s.Notifier.Notified == 1)
}
