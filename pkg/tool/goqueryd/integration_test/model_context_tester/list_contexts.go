package model_context_tester

import "github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"

func (t *Tester) ListContexts() string {
	return t.MustCallTool(constant.ListContexts, map[string]any{})
}
