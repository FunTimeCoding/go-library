package coordination

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/base"
	"testing"
	"time"
)

func TestInactivityTimeout(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "working on something")
	s.Advance(2 * time.Hour)
	s.Service.RunTimeoutSweep()
	r := a.CheckLive()
	assert.True(t, r.TimeoutMessage != nil)
	assert.StringContains(t, "inactivity", *r.TimeoutMessage)
}

func TestCompleteTimeout(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "a task")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "done",
		},
	)
	s.Advance(45 * time.Minute)
	s.Service.RunTimeoutSweep()
	r := a.CheckLive()
	assert.True(t, r.TimeoutMessage != nil)
	assert.StringContains(t, "completing", *r.TimeoutMessage)
}

func TestNoFalseTimeout(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "still working")
	s.Advance(30 * time.Minute)
	s.Service.RunTimeoutSweep()
	r := a.CheckLive()
	assert.True(t, r.TimeoutMessage == nil)
}

func TestTimeoutClearsOnCheck(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "working")
	s.Advance(2 * time.Hour)
	s.Service.RunTimeoutSweep()
	first := a.CheckLive()
	assert.True(t, first.TimeoutMessage != nil)
	second := a.CheckLive()
	assert.True(t, second.TimeoutMessage == nil)
}

func TestTimeoutVisibleToOthers(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	b := s.NewSession(t)
	defer b.Close()
	a.Announce(a.Name(), "going quiet")
	b.Announce(b.Name(), "watching")
	s.Advance(2 * time.Hour)
	s.Service.RunTimeoutSweep()
	v := b.MustCallTool(constant.History, map[string]any{})
	assert.StringContains(t, "timed out", v)
	assert.StringContains(t, a.Name(), v)
}
