package store_tester

import "github.com/funtimecoding/go-library/pkg/assert"

func (o *Tester) ConsumeTimeout(name string) string {
	result, e := o.Store.ConsumeTimeout(name)
	assert.FatalOnError(o.t, e)

	return result
}
