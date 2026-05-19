package coordination

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestRelease(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "temporary work")
	a.MustCallTool(constant.Release, map[string]any{})
	r := a.Check()
	assert.Count(t, 0, r.Sessions)
}

func TestReleaseHistoryEvent(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	b := s.NewSession(t)
	defer b.Close()
	a.Announce(a.Name(), "observer")
	b.Announce(b.Name(), "departing")
	b.MustCallTool(constant.Release, map[string]any{})
	h := a.MustCallTool(constant.History, map[string]any{})
	assert.StringContains(t, b.Name(), h)
	assert.StringContains(t, "left", h)
}
