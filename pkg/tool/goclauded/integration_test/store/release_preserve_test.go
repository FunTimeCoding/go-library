//go:build local

package store

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/store_tester"
	"testing"
)

func TestReleasePreservesSession(t *testing.T) {
	s := store_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Store.ReleaseByCallsign(r.Callsign)
	e := s.Store.GetSession("session-1")
	assert.True(t, e != nil)
	assert.String(t, r.Callsign, e.Name)
	assert.True(t, e.Callsign == nil)
}

func TestReleaseFreesCallsign(t *testing.T) {
	s := store_tester.New(t)
	var callsigns []string

	for i := 0; ; i++ {
		r := s.Store.EnsureSession(fmt.Sprintf("fill-%d", i))

		if r.Callsign == "" {
			break
		}

		callsigns = append(callsigns, r.Callsign)
	}

	released := callsigns[0]
	s.Store.ReleaseByCallsign(released)
	r := s.Store.EnsureSession("after-release")
	assert.String(t, released, r.Callsign)
}
