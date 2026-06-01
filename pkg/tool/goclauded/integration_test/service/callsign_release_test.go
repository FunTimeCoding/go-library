//go:build local

package service

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/service_tester"
	"sort"
	"testing"
	"time"
)

func TestCallsignReleasedAfterThreeDays(t *testing.T) {
	s := service_tester.New(t)
	r := s.Service.Register("session-1")
	assert.True(t, r.Callsign != "")
	callsign := r.Callsign
	s.Advance(73 * time.Hour)
	s.Service.RunTimeoutSweep()
	e := s.Store.GetSession("session-1")
	assert.True(t, e.Callsign == nil)
	assert.String(t, callsign, e.Name)
}

func TestCallsignNotReleasedBeforeThreeDays(t *testing.T) {
	s := service_tester.New(t)
	r := s.Service.Register("session-1")
	assert.True(t, r.Callsign != "")
	s.Advance(71 * time.Hour)
	s.Service.RunTimeoutSweep()
	e := s.Store.GetSession("session-1")
	assert.True(t, e.Callsign != nil)
}

func TestCallsignRecycledAfterSweepRelease(t *testing.T) {
	s := service_tester.New(t)
	expected := []string{
		"Ash", "Blair", "Cedar", "Dale", "Ellis",
		"Frost", "Glen", "Harbor", "Jade", "Kent",
	}
	var callsigns []string

	for i := 0; ; i++ {
		r := s.Service.Register(fmt.Sprintf("fill-%d", i))

		if r.Callsign == "" {
			break
		}

		callsigns = append(callsigns, r.Callsign)
	}

	sort.Strings(callsigns)
	assert.Integer(t, len(expected), len(callsigns))

	for i, v := range expected {
		assert.String(t, v, callsigns[i])
	}

	s.Advance(73 * time.Hour)
	s.Service.RunTimeoutSweep()
	r := s.Service.Register("after-sweep")
	assert.True(t, r.Callsign != "")
}
