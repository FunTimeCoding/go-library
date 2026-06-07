package store_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/event"
	"time"
)

func (o *Tester) EventsSince(
	since time.Time,
	before time.Time,
	kind string,
	limit int,
	offset int,
) []event.Event {
	result, e := o.Store.EventsSince(since, before, kind, limit, offset)
	assert.FatalOnError(o.t, e)

	return result
}
