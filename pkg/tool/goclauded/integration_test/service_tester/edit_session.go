package service_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service/argument/edit_session"
)

func (o *Tester) EditSession(
	identifier string,
	a *edit_session.Session,
) {
	assert.FatalOnError(o.t, o.Service.EditSession(identifier, a))
}
