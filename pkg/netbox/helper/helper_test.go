package helper

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	webConstant "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"testing"
)

func TestHelper(t *testing.T) {
	assert.String(
		t,
		"http://localhost/a",
		ToWebLink(
			locator.New(
				webConstant.Localhost,
			).Insecure().Base(constant.Interface).Path("/a").String(),
		),
	)
}
