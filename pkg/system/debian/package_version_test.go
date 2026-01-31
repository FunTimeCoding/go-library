package debian

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"testing"
)

func TestPackageVersion(t *testing.T) {
	assert.String(
		t,
		"example_1.0.0-1_amd64",
		PackageVersion(
			"example",
			"1.0.0",
			1,
			constant.AMD64,
		),
	)
}
