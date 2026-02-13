package reporter

import "github.com/funtimecoding/go-library/pkg/errors/sentry"

func (r *Reporter) Stop() {
	sentry.Flush(r.hub)
}
