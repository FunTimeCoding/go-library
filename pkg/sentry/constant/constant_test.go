package constant

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "SENTRY_LOCATOR", LocatorEnvironment)
	assert.String(t, "SENTRY_ORGANIZATION", OrganizationEnvironment)
	assert.String(t, "SENTRY_PROJECT", ProjectEnvironment)
}
