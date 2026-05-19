package coordination

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/base"
	"testing"
	"time"
)

func TestEdit(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "some work")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "hasty summary",
		},
	)
	events := s.Store().EventsSince(time.Time{}, time.Time{}, "", 1, 0)
	assert.Count(t, 1, events)
	a.MustCallTool(
		constant.Edit,
		map[string]any{
			constant.Identifier: float64(events[0].Identifier),
			constant.Message:    "corrected summary with design details",
		},
	)
	assert.StringContains(
		t,
		"corrected summary with design details",
		a.MustCallTool(constant.History, map[string]any{}),
	)
}

func TestEditNegativeID(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "bind identity")
	assert.StringContains(
		t,
		"invalid event identifier",
		a.MustCallToolError(
			constant.Edit,
			map[string]any{
				constant.Identifier: float64(-1),
				constant.Message:    "should fail",
			},
		),
	)
}
