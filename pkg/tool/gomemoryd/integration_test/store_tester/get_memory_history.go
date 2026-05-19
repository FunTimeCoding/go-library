package store_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
)

func (o *Tester) GetMemoryHistory(identifier int64) []store.Version {
	o.t.Helper()
	result, e := o.Store.GetMemoryHistory(identifier)
	assert.FatalOnError(o.t, e)

	return result
}
