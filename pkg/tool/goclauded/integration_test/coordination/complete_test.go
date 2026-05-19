package coordination

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestComplete(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building search index")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "search index and refresh implemented",
		},
	)
	r := a.Check()
	assert.Count(t, 1, r.Sessions)
	assert.String(t, "", r.Sessions[0].Topic)
}

func TestCompleteHistoryEvent(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "reviewing proposals")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "reviewed 8 proposals into design docs",
		},
	)
	history := a.MustCallTool(constant.History, map[string]any{})
	assert.StringContains(t, "completed reviewing proposals", history)
	assert.StringContains(t, "reviewed 8 proposals into design docs", history)
}

func TestCompleteAndReannounce(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "first task")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "done with first",
		},
	)
	a.Announce(a.Name(), "second task")
	check := a.Check()
	assert.String(t, "second task", check.Sessions[0].Topic)
}

func TestCompleteWithExplicitTopic(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "main work")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "main work done",
		},
	)
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Topic:   "ad-hoc fix",
			constant.Message: "fixed the spacing bug",
		},
	)
	history := a.MustCallTool(constant.History, map[string]any{})
	assert.StringContains(t, "completed ad-hoc fix", history)
	assert.StringContains(t, "fixed the spacing bug", history)
}

func TestCompleteNoTopicNoAnnounce(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "work")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "done",
		},
	)
	result := a.MustCallToolError(
		constant.Complete,
		map[string]any{
			constant.Message: "nothing to complete",
		},
	)
	assert.StringContains(t, "no active topic", result)
}

func TestCompleteBeforeAnnounce(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	result := a.MustCallToolError(
		constant.Complete,
		map[string]any{
			constant.Message: "nothing to complete",
		},
	)
	assert.StringContains(t, "announce first", result)
}
