package model_context_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert/fixture"
	system "github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"
)

func (t *Tester) IndexFixtures() {
	t.MustCallTool(
		constant.AddCollection,
		map[string]any{
			"name":       "test",
			constant.Path: fixture.Path(system.SearchPath),
		},
	)
	t.MustCallTool(constant.Index, map[string]any{})
}
