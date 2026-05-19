package store_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
)

func (o *Tester) ListTags() []store.TagCount {
	o.t.Helper()
	result, e := o.Store.ListTags()
	assert.FatalOnError(o.t, e)

	return result
}
