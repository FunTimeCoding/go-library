package store_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
)

func (o *Tester) ListRelations(identifier int64) []store.Relation {
	o.t.Helper()
	result, e := o.Store.ListRelations(identifier)
	assert.FatalOnError(o.t, e)

	return result
}
