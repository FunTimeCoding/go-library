//go:build local

package service

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	s := service_tester.New(t)
	result := s.Service.Register("session-1")
	assert.True(t, result.New)
	assert.True(t, result.Callsign != "")
	assert.True(t, s.Store.GetSession("session-1").NeedsRoster)
}

func TestRegisterExisting(t *testing.T) {
	s := service_tester.New(t)
	first := s.Service.Register("session-1")
	second := s.Service.Register("session-1")
	assert.True(t, first.New)
	assert.False(t, second.New)
	assert.String(t, first.Callsign, second.Callsign)
}

func TestRegisterLogsEvent(t *testing.T) {
	s := service_tester.New(t)
	s.Service.Register("session-1")
	assert.Count(
		t,
		1,
		s.Store.EventsSince(
			time.Time{},
			time.Time{},
			"register",
			10,
			0,
		),
	)
}

func TestRegisterDifferentNames(t *testing.T) {
	s := service_tester.New(t)
	r1 := s.Service.Register("session-1")
	r2 := s.Service.Register("session-2")
	assert.True(t, r1.Callsign != r2.Callsign)
}

func TestRegisterIncrementsTurnCount(t *testing.T) {
	s := service_tester.New(t)
	s.Service.Register("session-1")
	s.Service.Register("session-1")
	s.Service.Register("session-1")
	assert.True(t, s.Store.GetSession("session-1").TurnCount >= 2)
}
