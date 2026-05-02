package reporter

import "github.com/funtimecoding/go-library/pkg/errors/sentry"

func (r *Reporter) Start() {
	if r.locator == "" {
		return
	}

	r.hub = sentry.Start(
		r.project,
		r.locator,
		r.environment,
		r.version,
	)
}
