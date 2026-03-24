package reporter

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry"
	"os"
)

func (r *Reporter) RecoverFlush(v any) {
	r.hub.Recover(v)
	sentry.Flush(r.hub)

	if v != nil {
		os.Exit(1)
	}
}
