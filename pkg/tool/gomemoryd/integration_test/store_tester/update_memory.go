package store_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
)

func (o *Tester) UpdateMemory(
	identifier int64,
	option *store.SaveOption,
) {
	o.t.Helper()
	assert.FatalOnError(o.t, o.Store.UpdateMemory(identifier, option))
}
