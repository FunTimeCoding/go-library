//go:build local

package service

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
	"time"
)

func TestMoment(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Moment("session-1", r.Callsign, "figured out the root cause")
	events := s.Store.EventsSince(
		time.Time{},
		time.Time{},
		constant.Moment,
		10,
		0,
	)
	assert.Count(t, 1, events)
	assert.String(t, "figured out the root cause", events[0].Body)
}
