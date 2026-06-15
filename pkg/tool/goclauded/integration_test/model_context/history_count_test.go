package model_context

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestHistoryCountEmpty(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "bind identity")
	result := a.MustCallTool(constant.HistoryCount, map[string]any{})
	assert.StringContains(t, "1 events", result)
}

func TestHistoryCountMultiple(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "one")
	a.Announce(a.Name(), "two")
	a.MustCallTool(
		constant.Update,
		map[string]any{
			constant.Message: "completed",
			constant.Topic:   "three",
		},
	)
	result := a.MustCallTool(constant.HistoryCount, map[string]any{})
	assert.StringContains(t, "3 events", result)
}
