package store_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
)

func (o *Tester) CreateMemory(p *store.SaveOption) int64 {
	o.t.Helper()
	result, e := o.Store.CreateMemory(p)
	assert.FatalOnError(o.t, e)

	return result
}
