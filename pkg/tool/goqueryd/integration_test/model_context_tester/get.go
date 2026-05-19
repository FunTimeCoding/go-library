package model_context_tester

import "github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"

func (t *Tester) Get(path string) string {
	return t.MustCallTool(
		constant.Get,
		map[string]any{constant.Path: path},
	)
}
