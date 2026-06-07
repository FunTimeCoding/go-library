package service_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/ensure_result"
)

func (o *Tester) Register(sessionIdentifier string) *ensure_result.Result {
	result, e := o.Service.Register(sessionIdentifier)
	assert.FatalOnError(o.t, e)

	return result
}
