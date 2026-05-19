package model_context

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestRenameSelf(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "working")
	a.MustCallTool(
		constant.Rename,
		map[string]any{
			constant.Alias: "my-project",
		},
	)
	roster := a.MustCallTool(constant.Roster, map[string]any{})
	assert.StringContains(t, "my-project", roster)
}

func TestRenameOther(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	b := s.NewSession(t)
	defer b.Close()
	a.Announce(a.Name(), "renamer")
	b.Announce(b.Name(), "rename candidate")
	a.MustCallTool(
		constant.Rename,
		map[string]any{
			constant.Alias:  "the-target",
			constant.Target: b.Name(),
		},
	)
	roster := a.MustCallTool(constant.Roster, map[string]any{})
	assert.StringContains(t, "the-target", roster)
}
