package reporter

import "github.com/funtimecoding/go-library/pkg/errors/sentry"

func (r *Reporter) RecoverFlush(v any) {
	r.hub.Recover(v)
	sentry.Flush(r.hub)
}
