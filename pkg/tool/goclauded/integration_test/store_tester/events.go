package store_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/event_query"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/event"
)

func (o *Tester) Events(q *event_query.Query) []event.Event {
	result, e := o.Store.Events(q)
	assert.FatalOnError(o.t, e)

	return result
}
