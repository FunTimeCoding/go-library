package debian

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/constant"
	system "github.com/funtimecoding/go-library/pkg/system/constant"
	"testing"
)

func TestPackageVersion(t *testing.T) {
	assert.String(
		t,
		"example_1.0.0-1_amd64",
		PackageVersion(
			"example",
			constant.DefaultVersion,
			1,
			system.AMD64,
		),
	)
}
