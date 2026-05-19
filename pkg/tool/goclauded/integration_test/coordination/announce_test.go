package coordination

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestAnnounce(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "reviewing proposals")
	r := a.Check()
	assert.Count(t, 1, r.Sessions)
	assert.String(t, a.Name(), r.Sessions[0].Name)
	assert.String(t, "reviewing proposals", r.Sessions[0].Topic)
}

func TestAnnounceReannounce(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "first topic")
	a.Announce(a.Name(), "second topic")
	r := a.Check()
	assert.Count(t, 1, r.Sessions)
	assert.String(t, "second topic", r.Sessions[0].Topic)
}
