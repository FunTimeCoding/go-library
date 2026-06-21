package store_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store/save_option"
)

func (o *Tester) UpdateMemory(
	identifier int64,
	option *save_option.Option,
) {
	o.t.Helper()
	assert.FatalOnError(o.t, o.Store.UpdateMemory(identifier, option))
}
