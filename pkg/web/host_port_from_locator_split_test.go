package web

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"testing"
)

func TestHostPortFromLocatorSplit(t *testing.T) {
	host, port := HostPortFromLocatorSplit(
		locator.New(constant.Example).Port(8080).Path("/a").String(),
	)
	assert.String(t, "example.org", host)
	assert.Integer(t, 8080, port)
}
