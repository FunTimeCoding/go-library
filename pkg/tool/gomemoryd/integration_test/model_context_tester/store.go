package model_context_tester

import "github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"

func (t *Tester) Store() *store.Store {
	return t.base.Store()
}
