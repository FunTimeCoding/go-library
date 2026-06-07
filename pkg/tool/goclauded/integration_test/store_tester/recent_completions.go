package store_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/completion"
)

func (o *Tester) RecentCompletions() []completion.Completion {
	result, e := o.Store.RecentCompletions()
	assert.FatalOnError(o.t, e)

	return result
}
