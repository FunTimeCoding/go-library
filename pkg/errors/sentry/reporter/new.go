package reporter

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func New(
	project string,
	version string,
) *Reporter {
	return &Reporter{
		project: project,
		locator: environment.Required(constant.LocatorEnvironment),
		version: version,
	}
}
