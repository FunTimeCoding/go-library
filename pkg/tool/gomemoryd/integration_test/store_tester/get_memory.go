package store_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
)

func (o *Tester) GetMemory(identifier int64) *store.Memory {
	o.t.Helper()
	m, e := o.Store.GetMemory(identifier)
	assert.FatalOnError(o.t, e)

	return m
}
