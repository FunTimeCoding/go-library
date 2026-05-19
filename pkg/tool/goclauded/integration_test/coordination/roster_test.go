package coordination

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestRosterTwoSessions(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	b := s.NewSession(t)
	defer b.Close()
	a.Announce(a.Name(), "search index")
	b.Announce(b.Name(), "schema migration")
	roster := a.MustCallTool(constant.Roster, map[string]any{})
	assert.StringContains(t, "search index", roster)
	assert.StringContains(t, "schema migration", roster)
}

func TestRosterAfterRelease(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	b := s.NewSession(t)
	defer b.Close()
	a.Announce(a.Name(), "staying")
	b.Announce(b.Name(), "leaving")
	b.MustCallTool(constant.Release, map[string]any{})
	r := a.MustCallTool(constant.Roster, map[string]any{})
	assert.StringContains(t, "staying", r)
	assert.StringNotContains(t, "leaving", r)
}

func TestRosterAfterComplete(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "some task")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "done",
		},
	)
	r := a.Check()
	assert.Count(t, 1, r.Sessions)
	assert.String(t, "", r.Sessions[0].Topic)
}
