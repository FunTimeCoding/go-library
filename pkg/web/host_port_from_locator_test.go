package web

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"testing"
)

func TestHostPortFromLocator(t *testing.T) {
	assert.String(
		t,
		"example.org:8080",
		HostPortFromLocator(
			locator.New(
				constant.Example,
			).Port(8080).Path("/a").String(),
		),
	)
}
