package web

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"testing"
)

func TestParseLocator(t *testing.T) {
	actual := ParseLocator(
		locator.New(constant.Example).Port(8080).String(),
	)
	assert.String(t, "https://example.org:8080", actual.String())
	assert.String(t, "https", actual.Scheme)
	assert.String(t, "example.org:8080", actual.Host)
	assert.String(t, "8080", actual.Port())
	assert.String(t, "", actual.Path)
	assert.String(t, "", actual.RawQuery)
}
