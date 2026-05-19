package coordination

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestReannounceAfterRestart(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building things")
	s.Store().ClearBindings()
	r := a.CheckLive()
	assert.True(t, r.Reannounce != nil && *r.Reannounce)
}

func TestReannounceConsumedOnCheck(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building things")
	s.Store().ClearBindings()
	first := a.CheckLive()
	assert.True(t, first.Reannounce != nil && *first.Reannounce)
	second := a.CheckLive()
	assert.True(t, second.Reannounce == nil)
}

func TestReannounceNotSetWithoutPriorBinding(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	s.Store().ClearBindings()
	r := a.CheckLive()
	assert.True(t, r.Reannounce == nil)
}

func TestAnnounceClearsReannounce(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "first topic")
	s.Store().ClearBindings()
	a.Announce(a.Name(), "second topic")
	r := a.CheckLive()
	assert.True(t, r.Reannounce == nil)
}
