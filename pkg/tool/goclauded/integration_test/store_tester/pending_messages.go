package store_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/message"
)

func (o *Tester) PendingMessages(name string) []message.Message {
	result, e := o.Store.PendingMessages(name)
	assert.FatalOnError(o.t, e)

	return result
}
