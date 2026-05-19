package model_context_tester

import "github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"

func (t *Tester) Delete(collection string, path string) string {
	return t.MustCallTool(
		constant.Delete,
		map[string]any{
			constant.Collection: collection,
			constant.Path:       path,
		},
	)
}
