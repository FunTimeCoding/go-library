package coordination

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestUpdate(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "first topic")
	a.MustCallTool(
		constant.Update,
		map[string]any{
			constant.Message: "completed",
			constant.Topic:   "second topic",
		},
	)
	r := a.Check()
	assert.Count(t, 1, r.Sessions)
	assert.String(t, "second topic", r.Sessions[0].Topic)
}

func TestUpdateMultiple(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "step one")
	a.MustCallTool(
		constant.Update,
		map[string]any{
			constant.Message: "completed",
			constant.Topic:   "step two",
		},
	)
	a.MustCallTool(
		constant.Update,
		map[string]any{
			constant.Message: "completed",
			constant.Topic:   "step three",
		},
	)
	r := a.Check()
	assert.String(t, "step three", r.Sessions[0].Topic)
}

func TestUpdateHistoryEvents(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "initial")
	a.MustCallTool(
		constant.Update,
		map[string]any{
			constant.Message: "completed",
			constant.Topic:   "milestone one",
		},
	)
	r := a.MustCallTool(constant.History, map[string]any{})
	assert.StringContains(t, "updated milestone one: completed", r)
}

func TestUpdateBeforeAnnounce(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	assert.StringContains(
		t,
		"announce first",
		a.MustCallToolError(
			constant.Update,
			map[string]any{
				constant.Message: "completed",
				constant.Topic:   "premature update",
			},
		),
	)
}
