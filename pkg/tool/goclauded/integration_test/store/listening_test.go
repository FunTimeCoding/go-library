//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/store_tester"
	"testing"
)

func TestSetListening(t *testing.T) {
	s := store_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Store.SetListening(r.Callsign, true)
	e := s.Store.GetSession("session-1")
	assert.True(t, e.Listening)
	s.Store.SetListening(r.Callsign, false)
	e = s.Store.GetSession("session-1")
	assert.False(t, e.Listening)
}
