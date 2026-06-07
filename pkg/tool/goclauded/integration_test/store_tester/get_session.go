package store_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (o *Tester) GetSession(sessionIdentifier string) *session.Session {
	result, e := o.Store.GetSession(sessionIdentifier)
	assert.FatalOnError(o.t, e)

	return result
}
