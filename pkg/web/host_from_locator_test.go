package web

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"testing"
)

func TestHostFromLocator(t *testing.T) {
	assert.String(
		t,
		"example.org",
		HostFromLocator(
			locator.New(constant.Example).Port(8080).Path("/a").String(),
		),
	)
	assert.String(
		t,
		"example.org",
		HostFromLocator(locator.New(constant.Example).Path("/a").String()),
	)
}
