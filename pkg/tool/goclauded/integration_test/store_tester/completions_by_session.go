package store_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/completion"
)

func (o *Tester) CompletionsBySession(
	sessionIdentifier string,
) []completion.Completion {
	result, e := o.Store.CompletionsBySession(sessionIdentifier)
	assert.FatalOnError(o.t, e)

	return result
}
