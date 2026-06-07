package store_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (o *Tester) ListSessions() []session.Session {
	result, e := o.Store.ListSessions()
	assert.FatalOnError(o.t, e)

	return result
}
