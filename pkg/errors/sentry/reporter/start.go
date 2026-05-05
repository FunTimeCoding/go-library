package reporter

import "github.com/funtimecoding/go-library/pkg/errors/sentry"

func (r *Reporter) Start() *Reporter {
	if r.locator == "" {
		return nil
	}

	r.hub = sentry.Start(r.project, r.locator, r.version)

	return r
}
