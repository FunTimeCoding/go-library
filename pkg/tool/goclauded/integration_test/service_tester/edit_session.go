package service_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service/argument"
)

func (o *Tester) EditSession(
	identifier string,
	a *argument.EditSession,
) {
	assert.FatalOnError(o.t, o.Service.EditSession(identifier, a))
}
